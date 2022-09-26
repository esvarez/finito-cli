package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/esvarez/finito/internal/entity"
	"google.golang.org/api/sheets/v4"
)

type sheetRepo interface {
	Create(ctx context.Context, nameSpreadSheet, mainSheet string) (string, error)
	AddSheet(ctx context.Context, sheetID, name string) error
	AddRow(ctx context.Context, sheetID, column string, row *entity.Transaction) error
	ReadRange(ctx context.Context, sheetID, rangeName string) (*sheets.ValueRange, error)
	WriteValues(ctx context.Context, sheetID string, data []*sheets.ValueRange) error
}

const (
	_mainSheet         = "Summary"
	_transactionsSheet = "Transacciones"
	_layout            = "02/01/2006 15:04"
	_columExpense      = "B"
	_columIncome       = "I"
)

type SheetUseCase struct {
	repo sheetRepo
}

func NewSheet(repo sheetRepo) *SheetUseCase {
	return &SheetUseCase{
		repo: repo,
	}
}

func (u *SheetUseCase) CreateFinito(ctx context.Context, name string) (string, error) {
	sheetID, err := u.repo.Create(ctx, name, _mainSheet)
	if err != nil {
		return "", fmt.Errorf("failed to create sheet: %w", err)
	}
	if err = u.repo.AddSheet(ctx, sheetID, "Transactions"); err != nil {
		return "", fmt.Errorf("failed to add sheet: %w", err)
	}
	return sheetID, nil
}

// Deprecated: use AddIncome instead
func (u *SheetUseCase) AddExpense(ctx context.Context, sheetID string, transaction *entity.Transaction) error {
	column := _transactionsSheet + "!"
	u.fillDefaultTransactionValues(transaction)
	return u.repo.AddRow(ctx, sheetID, column, transaction)
}

// Deprecated: use AddIncome instead
func (u *SheetUseCase) AddIncome(ctx context.Context, sheetID string, transaction *entity.Transaction) error {
	column := _transactionsSheet + "!"
	u.fillDefaultTransactionValues(transaction)
	return u.repo.AddRow(ctx, sheetID, column, transaction)
}

func (u *SheetUseCase) fillDefaultTransactionValues(transaction *entity.Transaction) {
	if v := transaction.Date; v == "" {
		transaction.Date = time.Now().Format("2006-01-02")
		transaction.Month = fmt.Sprintf("=MONTH(%s)", transaction.Date)
	}
	if v := transaction.Description; v == "" {
		transaction.Description = "//TODO"
	}
	if v := transaction.Category; v == "" {
		transaction.Category = "Otros"
	}
}

func (u *SheetUseCase) TransferExpenses(ctx context.Context, rowExpense, rowIncome int, sheetOrigin, SheetID string) error {
	resp, err := u.repo.ReadRange(ctx, sheetOrigin, "Wallet!A2:X100")
	if err != nil {
		return fmt.Errorf("failed to read range: %w", err)
	}
	expenseLayout := fmt.Sprintf("%s!%s", _transactionsSheet, _columExpense)
	incomeLayout := fmt.Sprintf("%s!%s", _transactionsSheet, _columIncome)

	data := make([]*sheets.ValueRange, 0)
	const transferencia = "Transferencia"

	for _, row := range resp.Values {
		if len(row) == 0 {
			break
		}
		t, err := time.Parse(_layout, row[0].(string))
		if err != nil {
			fmt.Println(err)
		}

		amount := fmt.Sprintf("%s", row[10])
		account := row[3].(string)
		if row[6] == "TRANSFER" {

			from := row[3].(string)
			to := row[9].(string)

			data = append(data, &sheets.ValueRange{
				Range: fmt.Sprintf("%s%d", expenseLayout, rowExpense),
				Values: [][]interface{}{
					{t.String(), int(t.Month()), amount, transferencia, transferencia, from},
				},
			})
			rowExpense++
			data = append(data, &sheets.ValueRange{
				Range: fmt.Sprintf("%s%d", incomeLayout, rowIncome),
				Values: [][]interface{}{
					{t.String(), int(t.Month()), amount, transferencia, transferencia, to},
				},
			})
			rowIncome++
		} else if row[1] == transferencia {
			description := row[1].(string)
			category := row[2].(string)

			data = append(data, &sheets.ValueRange{
				Range: fmt.Sprintf("%s%d", expenseLayout, rowExpense),
				Values: [][]interface{}{
					{t.String(), int(t.Month()), amount, description, category, account},
				},
			})
			rowExpense++

			data = append(data, &sheets.ValueRange{
				Range: fmt.Sprintf("%s%d", incomeLayout, rowIncome),
				Values: [][]interface{}{
					{t.String(), int(t.Month()), amount, description, transferencia, category},
				},
			})
			rowIncome++
		} else {
			description := row[1].(string)
			category := row[2].(string)

			value := [][]interface{}{
				{t.String(), int(t.Month()), amount, description, category, account},
			}

			if row[6] == "EXPENSE" {
				data = append(data, &sheets.ValueRange{
					Range:  fmt.Sprintf("%s%d", expenseLayout, rowExpense),
					Values: value,
				})
				rowExpense++
			} else if row[6] == "INCOME" {
				data = append(data, &sheets.ValueRange{
					Range:  fmt.Sprintf("%s%d", incomeLayout, rowIncome),
					Values: value,
				})
				rowIncome++
			}
		}
	}

	return u.repo.WriteValues(ctx, SheetID, data)
}
