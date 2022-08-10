package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finito",
	Short: "finito the command line app that helps you manage your finances",
	Long:  `finito is a command line app that helps you manage your finances.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(loginCmd())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
