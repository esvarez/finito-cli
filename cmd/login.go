package cmd

import (
	"github.com/esvarez/finito/pkg/google"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
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
