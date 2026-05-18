/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	getmetrics "github.com/Platon223/Larb/internal/commands/getMetrics"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var metricType string

// getMetricsCmd represents the getMetrics command
var getMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Gets the desired type of metrics.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.WithSpinner("Fetching metrics...", func() error {
			user := getmetrics.ConfigUser(viper.GetString("apiKey"))

			allowedTypes := map[string]struct{}{
				"count": {},
				"speed": {},
				"error": {},
			}

			if _, ok := allowedTypes[metricType]; !ok {
				fmt.Println("Allowed types are: count, speed, error.")
			}

			if metricType == "count" {

				metrics, err := user.GetLogCountMetric()

				if err != nil {
					fmt.Println(err)

					return err
				}

				fmt.Println("\n")
				fmt.Println("Log Count Per Day")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

				max := 0

				for _, service := range metrics {
					for _, day := range service.LogCounts {
						if day.Count > max {
							max = day.Count
						}
					}
				}

				for _, service := range metrics {
					fmt.Println(service.ServiceName)
					for i, day := range service.LogCounts {
						if i != 0 {
							filled := (day.Count * 20) / max
							empty := 20 - filled
							bars := strings.Repeat("█", filled) + strings.Repeat("░", empty)
							fmt.Printf("  %-12s %s %d\n", day.Date, bars, day.Count)

							if i != len(service.LogCounts)-1 {
								fmt.Println("\n")
							}
						}
					}

					fmt.Println()
				}
			}

			if metricType == "speed" {

				metrics, err := user.GetSpeedMetric()

				if err != nil {
					fmt.Println(err)

					return err
				}

				fmt.Println("\n")
				fmt.Println("Log Speed")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

				max := 0.0

				for _, service := range metrics {
					if service.Speed > max {
						max = service.Speed
					}
				}

				for _, service := range metrics {
					fmt.Println(service.ServiceName)
					fmt.Println("\n")
					filled := (service.Speed * 20) / max
					empty := 20 - filled
					bars := strings.Repeat("█", int(filled)) + strings.Repeat("░", int(empty))
					fmt.Printf("   %s %.1f %s\n", bars, service.Speed, "(Logs / Per Second)")

					fmt.Println()
				}
			}

			if metricType == "error" {

				metrics, err := user.GetErrorMetric()

				if err != nil {
					fmt.Println(err)

					return err
				}

				fmt.Println("\n")
				fmt.Println("Log Error Rate")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

				max := 0.0

				for _, service := range metrics {
					if service.Rate > max {
						max = service.Rate
					}
				}

				for _, service := range metrics {
					fmt.Println(service.ServiceName)
					fmt.Println("\n")
					fmt.Printf("   %.2f%s %s\n", service.Rate, "%", "(Percent of alert triggering logs)")

					fmt.Println()
				}
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(getMetricsCmd)

	getMetricsCmd.Flags().StringVarP(&metricType, "type", "t", "", "Metric Type")
	getMetricsCmd.MarkFlagRequired("type")
}
