/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Platon223/Larb/internal/domain/update"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string = "v1.0.8"

// Root command

var rootCmd = &cobra.Command{
	Use:   "larb",
	Short: "Larb. CLI that connects developers with LogArbor all through the terminal.",
	Long:  `Larb is a CLI tool that is used in order to interact with LogArbor. Larb allows developers to observe their applications' logs, alerts, and metrics in the terminal. Larb also has an extra feature that doesn't exist on the LogArbor platform itself: Logby, it is an AI chat assistant that helps developers get started with LogArbor and Larb. This tool is built for developers who like to do everything in their terminal even if it is Log Managment.`,
}

func Execute() {

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

// Init function

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Set up for viper

func initConfig() {

	home, _ := os.UserHomeDir()
	viper.AddConfigPath(home)
	viper.SetConfigName(".larb")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	// Checks the new release for every 24 hours
	lastCheck := viper.GetTime("lastUpdateCheck")
	if time.Since(lastCheck) >= 24*time.Hour {
		update.CheckReleases(Version)
		viper.Set("lastUpdateCheck", time.Now())
		viper.WriteConfig()
	}
}
