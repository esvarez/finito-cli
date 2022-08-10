package repo

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type SheetRepo struct {
	srv *sheets.Service
}

func NewSheetRepo(srv *sheets.Service) *SheetRepo {
	return &SheetRepo{
		srv: srv,
	}
}

func (r *SheetRepo) Create(ctx context.Context, name string) (string, error) {
	sheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: name,
		},
	}
	resp, err := r.srv.Spreadsheets.Create(sheet).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	return resp.SpreadsheetId, nil
}
