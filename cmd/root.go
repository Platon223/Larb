/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Root command

var rootCmd = &cobra.Command{
	Use:   "larb",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
}
