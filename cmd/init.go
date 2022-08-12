package cmd

import (
	"context"
	"log"

	"github.com/esvarez/finito/config"

	"github.com/spf13/cobra"
)

const (
	_nameSpreadsheet = "finito_test"
)

type initCmd struct {
	cfg   *config.Configuration
	sheet sheetUseCase
}

func newInitCmd(cfg *config.Configuration, sheet sheetUseCase) *initCmd {
	return &initCmd{
		cfg:   cfg,
		sheet: sheet,
	}
}

func (i initCmd) command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the project",
		Long:  `Initialize the project`,
		Run: func(cmd *cobra.Command, args []string) {
			if !isLoggedIn {
				log.Println("You must login first")
				return
			}
			if i.cfg.SheetID != nil {
				log.Println("Sheet already initialized")
				return
			}
			log.Println("Initializing the project")

			log.Println("Creating the spreadsheet")
			sheetID, err := i.sheet.CreateFinito(ctx, _nameSpreadsheet)
			if err != nil {
				log.Fatalf("Error creating spreadsheet: %v", err)
			}
			cfg := &config.Configuration{
				App: config.App{
					SheetID: &sheetID,
				},
			}

			log.Println("Saving the configuration")
			err = config.SaveConfiguration(cfg)
			if err != nil {
				log.Fatalf("Error saving configuration: %v", err)
			}

			log.Println("Project initialized")
		},
	}

	return cmd
}
