/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	allservices "github.com/Platon223/Larb/internal/commands/allServices"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// all command defined here

var allServicesCmd = &cobra.Command{
	Use:   "all",
	Short: "Gets all the services a user has created.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Runs with a spinner

		return ui.WithSpinner("Fetching services...", func() error {
			user := allservices.ConfigUser(viper.GetString("apiKey"))

			services, err := user.GetAllServices()

			if err != nil {
				fmt.Println()
				fmt.Println("Error: ", err)
				return nil
			}

			if len(services) == 0 {
				fmt.Println("")
				fmt.Println("No services Yet")
				fmt.Println("")
			} else {
				rows := "| Id | Name | Health | Alert Level | Total Logs |\n|---|---|---|---|---|\n"
				for _, service := range services {
					rows += fmt.Sprintf("| %s | %s | %s | %s | %d |\n",
						service.Id,
						service.Name,
						service.Health,
						service.AlertLevel,
						service.LogCount,
					)
				}

				r, _ := glamour.NewTermRenderer(
					glamour.WithStandardStyle("dracula"),
					glamour.WithWordWrap(0),
				)
				out, _ := r.Render(rows)
				fmt.Println(out)
			}
			return nil
		})

	},
}

// Inits the all command

func init() {
	rootCmd.AddCommand(allServicesCmd)
}
