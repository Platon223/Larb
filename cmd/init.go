/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	initCommd "github.com/Platon223/Larb/internal/commands/init"
	"github.com/spf13/cobra"
)

// init Command defined here

var initCmd = &cobra.Command{
	Use:   "init [apiKey]",
	Args:  cobra.ExactArgs(1),
	Short: "Init command connects the cli tool with your LogArbor account. The structure of init command: init [apiKey]",
	Run: func(cmd *cobra.Command, args []string) {
		result := initCommd.Init(args[0])

		fmt.Println("Status: ", result)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
