package chat

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
	openai "github.com/sashabaranov/go-openai"
)

func fetchDocs() (string, error) {

	// cache locations
	homeDir, _ := os.UserHomeDir()
	cacheFile := filepath.Join(homeDir, ".config", "larb", "docsCache.txt")
	metaFile := filepath.Join(homeDir, ".config", "larb", "docsCacheTime.txt")
	
	// check if the cache is fresh
	if data, err := os.ReadFile(metaFile); err == nil {
        lastFetch, _ := time.Parse(time.RFC3339, string(data))
        if time.Since(lastFetch) < 24*time.Hour {
            cached, _ := os.ReadFile(cacheFile)
			fmt.Println("Docs cached")
            return string(cached), nil
        }
    }

	larbReadme, err := http.Get("https://raw.githubusercontent.com/Platon223/Larb/main/README.md")
	if err != nil {
		return "", nil
	}

	logArborReadme, err := http.Get("https://raw.githubusercontent.com/Platon223/LogArbor/main/README.md")
	if err != nil {
		return "", err
	}
	defer logArborReadme.Body.Close()

	larbBytes, _ := io.ReadAll(larbReadme.Body)
	logArborBytes, _ := io.ReadAll(logArborReadme.Body)
	docs := fmt.Sprintf(`
	Your name is Logby and you are a AI Chat Assistant for Larb and LogArbor,
LARB CLI DOCS:
%s

LOGARBOR DOCS:
%s
    `, string(larbBytes), string(logArborBytes)) 

	os.MkdirAll(filepath.Dir(cacheFile), 0755)
    os.WriteFile(cacheFile, []byte(docs), 0644)
    os.WriteFile(metaFile, []byte(time.Now().Format(time.RFC3339)), 0644)

	return docs, nil
}

func StartChat(apiKey string) error {
	readme, err := fetchDocs()
	if err != nil {
		fmt.Println("Warning: could not fetch docs, using basic mode")
		readme = "Larb is a CLI tool for LogArbor log aggregation."
	}

	client := openai.NewClient(apiKey)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: readme,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	fig := figure.NewFigure("Logby", "doom", true)
	fig.Print()
	fmt.Println()
	out, _ := glamour.Render(`
# Logby

AI Chat Assistant. Ask me anything about Larb or LogArbor. Type 'exit' to quit

	`, "dracula")
	fmt.Println(out)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for {
		fmt.Print("\n> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "exit" || input == "quit" {
			fmt.Println("See you later!")
			return nil
		}

		if input == "" {
			continue
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT4oMini,
				Messages: messages,
			},
		)
		if err != nil {
			return fmt.Errorf("openai error: %w", err)
		}

		reply := resp.Choices[0].Message.Content

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: reply,
		})

		fmt.Println("\nLogby:", reply)
	}
}
