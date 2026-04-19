/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	initCommd "github.com/Platon223/Larb/internal/commands/init"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// init Command defined here

var initCmd = &cobra.Command{
	Use:   "init [apiKey]",
	Args:  cobra.ExactArgs(1),
	Short: "Init command connects the cli tool with your LogArbor account. The structure of init command: init [apiKey]",
	Run: func(cmd *cobra.Command, args []string) {
		initCommd.Init(args[0])

		fig := figure.NewFigure("LARB", "slant", true)
		fig.Print()

		fmt.Println()
		fmt.Println("Welcome to Larb — The LogArbor CLI")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		fmt.Println("Next Steps:")
		fmt.Println()
		fmt.Println("  1. Create a service")
		fmt.Println("     larb add --name Auth --level warning")
		fmt.Println()
		fmt.Println("  2. View your services")
		fmt.Println("     larb all")
		fmt.Println()
		fmt.Println("  3. Tail live logs")
		fmt.Println("     larb tail logs [serviceId]")
		fmt.Println()
		fmt.Println("  4. View metrics")
		fmt.Println("     larb metrics --type count")
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("For more commands go to the docs")
		fmt.Println("Docs: https://logarbor.com/docs")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
