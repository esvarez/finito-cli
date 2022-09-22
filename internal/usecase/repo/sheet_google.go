package repo

import (
	"context"

	"github.com/esvarez/finito/internal/entity"
	"google.golang.org/api/sheets/v4"
)

const (
	_overwrite   = "OVERWRITE"
	_userEntered = "USER_ENTERED"
)

type SheetRepo struct {
	srv *sheets.Service
}

func NewSheetRepo(srv *sheets.Service) *SheetRepo {
	return &SheetRepo{
		srv: srv,
	}
}

func (r *SheetRepo) Create(ctx context.Context, nameSpreadSheet, mainSheet string) (string, error) {
	sheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: nameSpreadSheet,
		},
		Sheets: []*sheets.Sheet{
			{
				Properties: &sheets.SheetProperties{Title: mainSheet},
			},
		},
	}
	resp, err := r.srv.Spreadsheets.Create(sheet).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	return resp.SpreadsheetId, nil
}

func (r *SheetRepo) AddRow(ctx context.Context, sheetID, column string, row *entity.Transaction) error {
	req := &sheets.ValueRange{
		Range: column,
		Values: [][]interface{}{
			{row.Date, row.Month, row.Amount, row.Description, row.Category, row.Account},
		},
	}

	_, err := r.srv.Spreadsheets.Values.
		Append(sheetID, column, req).
		ValueInputOption(_userEntered).
		InsertDataOption(_overwrite).
		Context(ctx).Do()
	if err != nil {
		return err
	}
	return nil
}

func (r *SheetRepo) AddSheet(ctx context.Context, sheetID, name string) error {
	requests := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title: name,
					},
				},
			},
		},
	}
	_, err := r.srv.Spreadsheets.BatchUpdate(sheetID, requests).Context(ctx).Do()
	if err != nil {
		return err
	}
	return nil
}

func (r *SheetRepo) ReadRange(ctx context.Context, sheetID, rangeName string) (*sheets.ValueRange, error) {
	resp, err := r.srv.Spreadsheets.Values.Get(sheetID, rangeName).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
