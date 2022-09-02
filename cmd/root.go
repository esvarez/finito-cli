package cmd

import (
	"github.com/esvarez/finito/pkg/ui"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "finito",
	Short: "A CLI for managing your finances",
}

func Execute() {

	controller := ui.NewController()

	addcmd := newAddCmd(controller)

	rootCmd.AddCommand(addcmd.command())

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
