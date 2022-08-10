package cmd

import "context"

type sheetUseCase interface {
	Create(ctx context.Context, name string) (string, error)
	AddExpense(ctx context.Context, sheetID string, expense string) error
	AddIncome(ctx context.Context, sheetID string, expense string) error
}
