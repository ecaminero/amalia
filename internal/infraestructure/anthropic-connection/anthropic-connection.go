package anthropic_connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const AnthropicAPIEndpoint = "https://api.anthropic.com/v1/complete"

type AnthropicConnection struct {
	Prompt        string   `json:"prompt"`
	Model         string   `json:"model"`
	MaxTokens     int      `json:"max_tokens_to_sample"`
	StopSequences []string `json:"stop_sequences,omitempty"`
	Temperature   float64  `json:"temperature,omitempty"`
	TopK          int      `json:"top_k,omitempty"`
	TopP          float64  `json:"top_p,omitempty"`
	ApiKey        string   `json:"api_key"`
}

type AnthropicResponse struct {
	Completion string `json:"completion"`
}

func NewAnthropicConnection() (*AnthropicConnection, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ANTHROPIC_API_KEY environment variable is not set")
	}

	model := os.Getenv("ANTHROPIC_MODEL")
	if model == "" {
		model = "claude-3-5-sonnet-20240620"
	}

	maxTokensStr := os.Getenv("ANTHROPIC_MAX_TOKENS")
	if maxTokensStr == "" {
		maxTokensStr = "1024"
	}

	maxTokens, err := strconv.Atoi(maxTokensStr)
	if err != nil {
		return nil, fmt.Errorf("error converting maxTokens to int: %w", err)
	}

	return &AnthropicConnection{
		ApiKey:      apiKey,
		Model:       model,
		MaxTokens:   maxTokens,
		Temperature: maxTokens || 0.7, // You can adjust these parameters
		TopK:        -1,               // as needed
		TopP:        -1,
	}, nil
}

func CallAnthropicAPI(receiver *AnthropicConnection, prompt string) (string, error) {

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %w", err)
	}

	req, err := http.NewRequest("POST", AnthropicAPIEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse AnthropicResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	return apiResponse.Completion, nil
}
