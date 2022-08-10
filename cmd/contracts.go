package cmd

import "context"

type sheetUseCase interface {
	Create(ctx context.Context, name string) (string, error)
}
