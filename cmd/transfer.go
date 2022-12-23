package cmd

import (
	"context"
	"fmt"
	"github.com/esvarez/finito/config"
	"github.com/esvarez/finito/internal/usecase"

	"github.com/spf13/cobra"
)

var (
	origin     string
	rowIncome  int
	rowExpense int
)

type Transfer struct {
	sheet *usecase.SheetUseCase
	conf  *config.Configuration
}

func NewCmdTransfer(sheet *usecase.SheetUseCase, configuration *config.Configuration) *Transfer {
	return &Transfer{
		sheet: sheet,
		conf:  configuration,
	}
}

func (t Transfer) CmdTransfer(ctx context.Context) *cobra.Command {
	var transferCmd = &cobra.Command{
		Use:   "transfer",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("transfer called")
			if err := t.sheet.TransferExpenses(ctx, origin, *t.conf.App.SheetID); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("transfer done")
		},
	}
	transferCmd.Flags().StringVarP(&origin, "origin", "o", "", "The origin sheet")

	transferCmd.MarkFlagRequired("origin")

	return transferCmd
}
