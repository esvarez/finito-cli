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

func (r *SheetRepo) AddExpense(ctx context.Context, sheetID string, expense string) error {
	field := "C3"
	return r.appendTransaction(ctx, sheetID, field, expense)
}

func (r *SheetRepo) AddIncome(ctx context.Context, sheetID string, expense string) error {
	field := "H3"
	return r.appendTransaction(ctx, sheetID, field, expense)
}

func (r *SheetRepo) appendTransaction(ctx context.Context, sheetID, column, expense string) error {
	req := &sheets.ValueRange{
		Values: [][]interface{}{
			{expense},
		},
	}

	valueImputOption := "USER_ENTERED"
	_, err := r.srv.Spreadsheets.Values.
		Append(sheetID, column, req).
		ValueInputOption(valueImputOption).
		Context(ctx).Do()

	if err != nil {
		return err
	}
	return nil
}
