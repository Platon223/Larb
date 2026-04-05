/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Platon223/Larb/internal/commands/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version of Larb",
	Run: func(cmd *cobra.Command, args []string) {
		version_string := version.Version()

		fmt.Println(version_string)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
