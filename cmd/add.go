package cmd

import (
	"context"
	"fmt"
	"github.com/esvarez/finito/internal/entity"
	"log"

	"github.com/esvarez/finito/config"

	"github.com/spf13/cobra"
)

var (
	amount      string
	date        string
	description string
	category    string
)

type addCmd struct {
	cfg   *config.App
	sheet sheetUseCase
}

func newAddCmd(cfg *config.App, sheet sheetUseCase) *addCmd {
	return &addCmd{
		cfg:   cfg,
		sheet: sheet,
	}
}

func (a addCmd) command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new transaction",
		Long:  `Add a new transaction`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
		},
	}

	cmd.AddCommand(a.addExpense(ctx))
	cmd.AddCommand(a.addIncome(ctx))

	return cmd
}

func (a addCmd) addExpense(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expense",
		Short: "Add a new expense",
		Long:  `Add a new expense`,
		Run: func(cmd *cobra.Command, args []string) {
			transaction := &entity.Transaction{
				Amount:      amount,
				Date:        date,
				Description: description,
				Category:    category,
			}
			err := a.sheet.AddExpense(ctx, *a.cfg.SheetID, transaction)
			if err != nil {
				log.Printf("Error adding expense: %v", err)
				return
			}
			fmt.Println("expense added")
		},
	}

	cmd.Flags().StringVarP(&amount, "amount", "a", "", "Amount of the expense")
	cmd.MarkFlagRequired("amount")
	return cmd
}

func (a addCmd) addIncome(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "income",
		Short: "Add a new income",
		Long:  `Add a new income`,
		Run: func(cmd *cobra.Command, args []string) {
			transaction := &entity.Transaction{
				Amount:      amount,
				Date:        date,
				Description: description,
				Category:    category,
			}
			err := a.sheet.AddIncome(ctx, *a.cfg.SheetID, transaction)
			if err != nil {
				log.Printf("Error adding income: %v", err)
				return
			}
			fmt.Println("income added")
		},
	}
	cmd.Flags().StringVarP(&amount, "amount", "a", "", "Amount of the expense")
	cmd.MarkFlagRequired("amount")
	return cmd
}
