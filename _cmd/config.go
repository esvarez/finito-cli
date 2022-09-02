package _cmd

import (
	"github.com/esvarez/finito/config"

	"github.com/spf13/cobra"
)

var (
	spreadSheetID string
)

func configCmd(cfg *config.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the project",
		Long:  "Configure the project",
		Run: func(cmd *cobra.Command, args []string) {
			cfg.SheetID = &spreadSheetID
			config.SaveConfiguration(cfg)
		},
	}

	cmd.Flags().StringVarP(&spreadSheetID, "sheet-id", "", "i", "The id of the spreadsheet")
	cmd.MarkFlagRequired("sheet-id")
	return cmd
}
