package cmd

import (
	"github.com/esvarez/finito/pkg/google"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to the server",
		Long:  "Login to the server",
		Run: func(cmd *cobra.Command, args []string) {
			google.Login()
		},
	}

	return cmd
}
