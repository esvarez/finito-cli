package usecase

import "context"

type sheetRepo interface {
	Create(context.Context, string) (string, error)
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
