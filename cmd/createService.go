/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	createservice "github.com/Platon223/Larb/internal/commands/createService"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string
var alertLevel string

// createServiceCmd represents the createService command
var createServiceCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Runs with a spinner

		return ui.WithSpinner("Creating your new service...", func() error {

			allowedLevels := map[string]struct{}{
				"debug": {},
				"info": {},
				"warning": {},
				"error": {},
				"critical": {},
			}

			if _, ok := allowedLevels[alertLevel]; !ok {
				fmt.Println("Allowed Alert Levels: [debug, info, warning, error, critical]")
				return nil
			}
			
			user := createservice.ConfigUser(viper.GetString("apiKey"))

			result, err := user.CreateService(name, alertLevel)

			if err != nil {
				fmt.Println()
				fmt.Println(err)
				return nil
			}

			fmt.Println()
			fmt.Println("Status: ", result)
			fmt.Println()

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(createServiceCmd)
	createServiceCmd.Flags().StringVarP(&name, "name", "n", "", "Service Name")
	createServiceCmd.Flags().StringVarP(&alertLevel, "level", "l", "", "Alert Level")
	createServiceCmd.MarkFlagRequired("name")
	createServiceCmd.MarkFlagRequired("level")
}
