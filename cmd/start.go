/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func startCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Long:  `Start the application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start called")
		},
	}
}
