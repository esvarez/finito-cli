package cmd

import (
	"context"
	"github.com/esvarez/finito/pkg/google"
	"log"
	"os"

	"github.com/esvarez/finito/internal/usecase"
	"github.com/esvarez/finito/internal/usecase/repo"

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
	ctx := context.Background()

	srv := google.GetService()
	if srv == nil {
		log.Println("You should login first to see all the options")
	}

	sheetRepo := repo.NewSheetRepo(srv)

	sheet := usecase.NewSheet(sheetRepo)

	rootCmd.AddCommand(loginCmd(srv))
	if srv != nil {
		rootCmd.AddCommand(initCmd(ctx, sheet))
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
