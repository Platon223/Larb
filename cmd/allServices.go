/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	allservices "github.com/Platon223/Larb/internal/commands/allServices"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// all command defined here

var allServicesCmd = &cobra.Command{
	Use:   "all",
	Short: "Gets all the services a user has created.",
	Run: func(cmd *cobra.Command, args []string) {

		user := allservices.ConfigUser(viper.GetString("apiKey"))

		services, err := user.GetAllServices()

		if err != nil {
			fmt.Println(err)
		}

		if len(services) == 0 {
			fmt.Println("\n")
			fmt.Println("No services Yet")
			fmt.Println("\n")
		} else {
			for i, service := range services {
				fmt.Println("\n")
				fmt.Println("Name: " + service.Name)
				fmt.Println("Id: " + service.Id)
				fmt.Println("Alert Level: " + service.AlertLevel)
				fmt.Println("Health: " + service.Health)
				fmt.Println("Total Logs: " + strconv.Itoa(service.LogCount))
				if i == len(services)-1 {
					fmt.Println("\n")
				}
			}
		}
	},
}

// Inits the all command

func init() {
	rootCmd.AddCommand(allServicesCmd)
}
