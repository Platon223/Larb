/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Platon223/Larb/internal/commands/version"
	"github.com/Platon223/Larb/internal/ui"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version of Larb",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ui.WithSpinner("Getting the version...", func() error {
			version_string := version.Version()

			fmt.Println(version_string)

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
