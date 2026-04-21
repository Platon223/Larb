package chat

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	openai "github.com/sashabaranov/go-openai"
)

func fetchDocs() (string, error) {
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

	return fmt.Sprintf(`
	Your name is Logby and you are a AI Chat Assistant for Larb and LogArbor,
LARB CLI DOCS:
%s

LOGARBOR DOCS:
%s
    `, string(larbBytes), string(logArborBytes)), nil
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

	fig := figure.NewFigure("LOGBY", "standard", true)
	fig.Print()
	fmt.Println()
	fmt.Println("Logby — AI Chat Assistant. Ask me anything about Larb or LogArbor (type 'exit' to quit)")
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
