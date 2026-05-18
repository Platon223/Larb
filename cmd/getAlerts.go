/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/charmbracelet/glamour"

	getalerts "github.com/Platon223/Larb/internal/commands/getAlerts"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getAlertsCmd represents the getAlerts command
var getAlertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "Gets user's alerts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.WithSpinner("Fetching alerts...", func() error {
			user := getalerts.ConfigUser(viper.GetString("apiKey"))

			alerts, err := user.GetAlerts()

			if err != nil {
				fmt.Println(err)
			}

			rows := "| Message | Level | Time | ServiceName | Viewed |\n|---|---|---|---|---|\n"

			for _, alert := range alerts {

				viewed := ""

				if alert.Viewed {
					viewed = "True"
				} else {
					viewed = "False"
				}

				rows += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
					alert.Message[:28]+"...",
					alert.Level,
					alert.Time,
					alert.ServiceName,
					viewed,
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
	rootCmd.AddCommand(getAlertsCmd)
}
