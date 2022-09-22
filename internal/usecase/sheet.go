package usecase

import (
	"context"
	"fmt"
	"google.golang.org/api/sheets/v4"
	"strings"
	"time"

	"github.com/esvarez/finito/internal/entity"
)

type sheetRepo interface {
	Create(ctx context.Context, nameSpreadSheet, mainSheet string) (string, error)
	AddSheet(ctx context.Context, sheetID, name string) error
	AddRow(ctx context.Context, sheetID, column string, row *entity.Transaction) error
	ReadRange(ctx context.Context, sheetID, rangeName string) (*sheets.ValueRange, error)
}

const (
	_mainSheet         = "Summary"
	_transactionsSheet = "Transacciones"
	_expenseField      = "B4"
	_incomeField       = "I4"
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

func (u *SheetUseCase) AddExpense(ctx context.Context, sheetID string, transaction *entity.Transaction) error {
	column := _transactionsSheet + "!" + _expenseField
	u.fillDefaultTransactionValues(transaction)
	return u.repo.AddRow(ctx, sheetID, column, transaction)
}

func (u *SheetUseCase) AddIncome(ctx context.Context, sheetID string, transaction *entity.Transaction) error {
	column := _transactionsSheet + "!" + _incomeField
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

func (u *SheetUseCase) TransferExpenses(ctx context.Context, sheetOrigin, SheetID string) error {

	resp, err := u.repo.ReadRange(ctx, sheetOrigin, "Wallet!A2:X100")
	if err != nil {
		return fmt.Errorf("failed to read range: %w", err)
	}

	for _, row := range resp.Values {
		if len(row) == 0 {
			break
		}
		date := strings.Split(row[0].(string), " ")[0]

		if row[6] == "TRANSFER" {

			tranferencia := "Transferencia"
			expense := &entity.Transaction{
				Date:        date,
				Month:       fmt.Sprintf("=MONTH(%s)", date),
				Amount:      fmt.Sprintf("=ABS(%s)", row[7]),
				Description: tranferencia,
				Category:    tranferencia,
				Account:     row[3].(string),
			}
			income := &entity.Transaction{
				Date:        date,
				Month:       fmt.Sprintf("=MONTH(%s)", date),
				Amount:      fmt.Sprintf("=ABS(%s)", row[7]),
				Description: tranferencia,
				Category:    tranferencia,
				Account:     row[9].(string),
			}

			if err := u.AddExpense(ctx, SheetID, expense); err != nil {
				return fmt.Errorf("failed to add expense: %w", err)
			}
			if err := u.AddIncome(ctx, SheetID, income); err != nil {
				return fmt.Errorf("failed to add income: %w", err)
			}
		} else if row[1] == "Ahorro" {
			if err := u.transferSavings(ctx, SheetID, date, row[4], row[1], row[2], row[3]); err != nil {
				return fmt.Errorf("failed to transfer savings: %w", err)
			}
		} else {
			date := strings.Split(row[0].(string), " ")[0]
			transaction := &entity.Transaction{
				Date:        date,
				Month:       fmt.Sprintf("=MONTH(%s)", date),
				Amount:      fmt.Sprintf("=ABS(%s)", row[4]),
				Description: row[1].(string),
				Category:    row[2].(string),
				Account:     row[3].(string),
			}

			if row[6] == "EXPENSE" {
				if err := u.AddExpense(ctx, SheetID, transaction); err != nil {
					return fmt.Errorf("failed to add expense: %w", err)
				}
			} else if row[6] == "INCOME" {
				if err := u.AddIncome(ctx, SheetID, transaction); err != nil {
					return fmt.Errorf("failed to add income: %w", err)
				}
			}
		}
	}

	return nil
}

func (u *SheetUseCase) transferSavings(ctx context.Context, id string, date, amount, description, category, account any) error {
	expense := &entity.Transaction{
		Date:        fmt.Sprintf("%s", date),
		Month:       fmt.Sprintf("=MONTH(%s)", date),
		Amount:      fmt.Sprintf("=ABS(%s)", amount),
		Description: fmt.Sprintf("%s", description),
		Category:    fmt.Sprintf("%s", category),
		Account:     fmt.Sprintf("%s", account),
	}

	income := &entity.Transaction{
		Date:        fmt.Sprintf("%s", date),
		Month:       fmt.Sprintf("=MONTH(%s)", date),
		Amount:      fmt.Sprintf("=ABS(%s)", amount),
		Description: fmt.Sprintf("%s", description),
		Category:    "Transferencia",
		Account:     fmt.Sprintf("%s", category),
	}

	if err := u.AddExpense(ctx, id, expense); err != nil {
		return fmt.Errorf("failed to add expense: %w", err)
	}
	if err := u.AddIncome(ctx, id, income); err != nil {
		return fmt.Errorf("failed to add income: %w", err)
	}
	return nil
}
