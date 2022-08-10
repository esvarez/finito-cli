package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
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
			//google.Login()
			fmt.Println(basepath)
			folder := "~/.finito/"
			if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
				err = os.Mkdir(folder, os.ModePerm)
				if err != nil {
					log.Printf("error creating folder %v", err)
				}
			}
		},
	}

	return cmd
}
