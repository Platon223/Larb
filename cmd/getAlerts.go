/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	getalerts "github.com/Platon223/Larb/internal/commands/getAlerts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getAlertsCmd represents the getAlerts command
var getAlertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "Gets user's alerts",
	Run: func(cmd *cobra.Command, args []string) {
		user := getalerts.ConfigUser(viper.GetString("apiKey"))

		alerts, err := user.GetAlerts()

		if err != nil {
			fmt.Println(err)
		}

		for i, alert := range alerts {
			fmt.Println("\n")
			fmt.Println("Id: ", alert.Id)
			fmt.Println("Message: ", alert.Message)
			fmt.Println("Level: ", strings.ToUpper(alert.Level))
			fmt.Println("Time: ", alert.Time)
			fmt.Println("UserId: ", alert.UserId)
			fmt.Println("ServiceId: ", alert.ServiceId)
			fmt.Println("Service Name: ", alert.ServiceName)
			if alert.Viewed {
				fmt.Println("Viewed: True")
			} else {
				fmt.Println("Viewed: False")
			}

			if i == len(alerts)-1 {
				fmt.Println("\n")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getAlertsCmd)
}
