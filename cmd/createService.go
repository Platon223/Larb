/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	createservice "github.com/Platon223/Larb/internal/commands/createService"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string
var alertLevel string

// createServiceCmd represents the createService command
var createServiceCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new service.",
	Run: func(cmd *cobra.Command, args []string) {
		user := createservice.ConfigUser(viper.GetString("apiKey"))

		result, err := user.CreateService(name, alertLevel)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Status: %s \n", result)
	},
}

func init() {
	rootCmd.AddCommand(createServiceCmd)
	createServiceCmd.Flags().StringVarP(&name, "name", "n", "", "Service Name")
	createServiceCmd.Flags().StringVarP(&alertLevel, "level", "l", "", "Alert Level")
	createServiceCmd.MarkFlagRequired("name")
	createServiceCmd.MarkFlagRequired("level")
}
