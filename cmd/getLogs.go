/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"

	getlogs "github.com/Platon223/Larb/internal/commands/getLogs"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var id string

// get command defined here
var getServiceCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the logs of a speciffic service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.WithSpinner("Fetching logs...", func() error {
			user := getlogs.ConfigUser(viper.GetString("apiKey"))

			logs, err := user.GetLogs(id)

			if err != nil {
				fmt.Println(err)
			}

			rows := "| Time | Level | Message |\n|---|---|---|\n"

			for _, log := range logs.LogList {

				rows += fmt.Sprintf("| %s | %s | %s |\n",
					log.Time,
					strings.ToUpper(log.Level),
					log.Message,
				)

			}

			r, _ := glamour.NewTermRenderer(
				glamour.WithStandardStyle("dracula"),
				glamour.WithWordWrap(0),
			)
			out, _ := r.Render(rows)
			fmt.Println(out)

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(getServiceCmd)
	getServiceCmd.Flags().StringVarP(&id, "id", "i", "", "Service ID")
	getServiceCmd.MarkFlagRequired("id")
}
