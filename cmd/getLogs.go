/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	getlogs "github.com/Platon223/Larb/internal/commands/getLogs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var id string

// get command defined here
var getServiceCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the logs of a speciffic service",
	Run: func(cmd *cobra.Command, args []string) {
		user := getlogs.ConfigUser(viper.GetString("apiKey"))

		logs, err := user.GetLogs(id)

		if err != nil {
			fmt.Println(err)
		}

		for i, log := range logs.LogList {
			if i == 0 {
				fmt.Println("\n")
			}

			logString := fmt.Sprintf("%s %s %s", log.Time, strings.ToUpper(log.Level), log.Message)
			fmt.Println(logString)
			if i == len(logs.LogList)-1 {
				fmt.Println("\n")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getServiceCmd)
	getServiceCmd.Flags().StringVarP(&id, "id", "i", "", "Service ID")
	getServiceCmd.MarkFlagRequired("id")
}
