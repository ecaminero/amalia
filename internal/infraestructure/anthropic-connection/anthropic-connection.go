package anthropic_connection

import (
	"context"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"os"
	"strconv"
)

type AnthropicConnection struct {
	Client       *anthropic.Client `json:"client"`
	Model        anthropic.Model   `json:"model"`
	SystemPrompt string            `json:"systemPrompt"`
	MaxTokens    int64             `json:"max_tokens"`
}

func NewAnthropicConnection() (*AnthropicConnection, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable is not set")
	}

	maxTokensStr := os.Getenv("ANTHROPIC_MAX_TOKENS")
	if maxTokensStr == "" {
		maxTokensStr = "1024"
	}

	modelKey := os.Getenv("ANTHROPIC_MODEL")
	model, found := GetModel(modelKey)
	if !found {
		return nil, fmt.Errorf("model not found: %s", modelKey)
	}

	maxTokens, err := strconv.ParseInt(maxTokensStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting maxTokens to int: %w", err)
	}

	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)
	var SystemPrompt = "Be very serious"

	return &AnthropicConnection{
		Client:       client,
		MaxTokens:    maxTokens,
		Model:        model,
		SystemPrompt: SystemPrompt,
	}, nil
}

func (a AnthropicConnection) GetComment(content string) (string, error) {
	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock(content)),
	}

	println("[user]: " + content)
	message, err := a.Client.Messages.New(context.TODO(),
		anthropic.MessageNewParams{
			Model: anthropic.F(a.Model),
			System: anthropic.F([]anthropic.TextBlockParam{
				anthropic.NewTextBlock(a.SystemPrompt),
			}), MaxTokens: anthropic.F(a.MaxTokens),
			Messages: anthropic.F(messages),
		})

	if err != nil {
		return "", err
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
	return message.Content[0].Text, nil

}
