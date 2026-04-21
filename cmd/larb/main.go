/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/Platon223/Larb/cmd"
	"github.com/joho/godotenv"
)

var OpenAIAPIKey = ""

func main() {
	godotenv.Load()
	os.Setenv("OPEN_AI_API_KEY", OpenAIAPIKey)
	cmd.Execute()
}
