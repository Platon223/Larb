/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	taillogs "github.com/Platon223/Larb/internal/commands/tailLogs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Tail command gets logs live.",
}

var logsLiveCmd = &cobra.Command{
	Use:   "logs [serviceId]",
	Short: "Stream live logs for a service.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		user := taillogs.ConfigUser(viper.GetString("apiKey"))

		return user.TailLogs(args[0])
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)

	tailCmd.AddCommand(logsLiveCmd)
}
