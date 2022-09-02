package _cmd

import (
	"context"

	"github.com/esvarez/finito/internal/entity"
)

type sheetUseCase interface {
	CreateFinito(ctx context.Context, name string) (string, error)
	AddExpense(ctx context.Context, sheetID string, transaction *entity.Transaction) error
	AddIncome(ctx context.Context, sheetID string, transaction *entity.Transaction) error
}

type viewController interface {
	Render() error
}
