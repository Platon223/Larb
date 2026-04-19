/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Platon223/Larb/internal/commands/chat"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat command lets you chat with Logby(AI Assistant).",
	RunE: func(cmd *cobra.Command, args []string) error {
		return chat.StartChat(os.Getenv("OPEN_AI_API_KEY"))
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
