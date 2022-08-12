package usecase

import (
	"context"
	"fmt"
	"github.com/esvarez/finito/internal/entity"
	"time"
)

type sheetRepo interface {
	Create(ctx context.Context, nameSpreadSheet, mainSheet string) (string, error)
	AddSheet(ctx context.Context, sheetID, name string) error
	AddRow(ctx context.Context, sheetID, column string, row *entity.Transaction) error
}

const (
	_mainSheet         = "Summary"
	_transactionsSheet = "Transacciones"
	_expenseField      = "B4"
	_incomeField       = "G4"
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
	}
	if v := transaction.Description; v == "" {
		transaction.Description = "//TODO"
	}
	if v := transaction.Category; v == "" {
		transaction.Category = "Other"
	}
}
