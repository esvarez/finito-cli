package _cmd

import (
	"context"

	"github.com/esvarez/finito/config"

	"github.com/spf13/cobra"
)

type startCmd struct {
	cfg  *config.App
	view viewController
}

func newStartCmd(cfg *config.App, view viewController) *startCmd {
	return &startCmd{
		cfg:  cfg,
		view: view,
	}
}

func (s startCmd) command(_ context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Long:  `Start the application`,
		Run: func(cmd *cobra.Command, args []string) {
			s.view.Render()
		},
	}

	return cmd
}
