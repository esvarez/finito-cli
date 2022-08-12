package cmd

import (
	"context"
	"github.com/esvarez/finito/config"
	"log"
	"os"

	"github.com/esvarez/finito/internal/usecase"
	"github.com/esvarez/finito/internal/usecase/repo"
	"github.com/esvarez/finito/pkg/google"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finito",
	Short: "finito the command line app that helps you manage your finances",
	Long:  `finito is a command line app that helps you manage your finances.`,
}

var isLoggedIn bool

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx := context.Background()
	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	srv := google.GetService()
	if srv != nil {
		isLoggedIn = true
	}

	sheetRepo := repo.NewSheetRepo(srv)

	sheet := usecase.NewSheet(sheetRepo)

	addCMD := newAddCmd(&cfg.App, sheet)
	initCMD := newInitCmd(cfg, sheet)

	rootCmd.AddCommand(loginCmd())
	rootCmd.AddCommand(configCmd(cfg))

	rootCmd.AddCommand(initCMD.command(ctx))
	rootCmd.AddCommand(addCMD.command(ctx))

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
