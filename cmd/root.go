package cmd

import (
	"context"
	"log"
	"os"

	"github.com/esvarez/finito/config"
	"github.com/esvarez/finito/internal/usecase"
	"github.com/esvarez/finito/internal/usecase/repo"
	"github.com/esvarez/finito/pkg/sheet"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "finito",
	Short: "A CLI for managing your finances",
}

func Execute() {
	ctx := context.Background()
	conf, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	srv, err := sheet.GetService()
	if err != nil {
		log.Fatalf("Error getting service: %v", err)
		return
	}

	sheetRepo := repo.NewSheetRepo(srv)
	sheet := usecase.NewSheet(sheetRepo)
	cmdTrnsfr := NewCmdTransfer(sheet, conf)

	rootCmd.AddCommand(cmdTrnsfr.CmdTransfer(ctx))
	rootCmd.AddCommand(loginCmd())

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
