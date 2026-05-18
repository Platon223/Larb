/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	initCommd "github.com/Platon223/Larb/internal/commands/init"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// init Command defined here

var initCmd = &cobra.Command{
	Use:   "init [apiKey]",
	Args:  cobra.ExactArgs(1),
	Short: "Init command connects the cli tool with your LogArbor account. The structure of init command: init [apiKey]",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.WithSpinner("Initializing the CLI...", func() error {
			result := initCommd.Init(args[0])

			if result != "Success" {
				fmt.Println()
				fmt.Println(result)
				return nil
			}

			fmt.Println()

			fig := figure.NewFigure("Larb", "doom", true)
			fig.Print()

			out, _ := glamour.Render(`

# Welcome to Larb

The LogArbor CLI

## Next Steps

1. Create a service: **larb add --name Auth --level warning**
2. View your services: **larb all**
3. Tail live logs: **larb tail logs [serviceId]**
4. View metrics: **larb metrics --type count**

## For more commands go to the docs

Docs: (https://logarbor.com/docs)
`, "dracula")

			fmt.Println(out)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
