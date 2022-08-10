package usecase

import "context"

type sheetRepo interface {
	Create(context.Context, string) (string, error)
	AddExpense(context.Context, string, string) error
	AddIncome(context.Context, string, string) error
}

type SheetUseCase struct {
	repo sheetRepo
}

func NewSheet(repo sheetRepo) *SheetUseCase {
	return &SheetUseCase{
		repo: repo,
	}
}

func (u *SheetUseCase) Create(ctx context.Context, name string) (string, error) {
	return u.repo.Create(ctx, name)
}

func (u *SheetUseCase) AddExpense(ctx context.Context, sheetID string, expense string) error {
	return u.repo.AddExpense(ctx, sheetID, expense)
}

func (u *SheetUseCase) AddIncome(ctx context.Context, sheetID string, expense string) error {
	return u.repo.AddIncome(ctx, sheetID, expense)
}
