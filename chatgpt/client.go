package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Req struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Rsp struct {
	ID      string         `json:"id"`
	Created int            `json:"created"`
	Model   string         `json:"model"`
	Choices []Choice       `json:"choices"`
	Usage   Usage          `json:"usage,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"` // Add an error field to capture error responses
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *string `json:"logprobs,omitempty"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ErrorResponse captures details of the error from API response
type ErrorResponse struct {
	Message string  `json:"message"`
	Type    string  `json:"type"`
	Param   *string `json:"param,omitempty"`
	Code    string  `json:"code"`
}

const (
	model3dot5Turbo    = "gpt-3.5-turbo"
	roleUser           = "user"
	roleAssistant      = "assistant"
	FinishReasonLength = "length"
)

func SendPrompt(prompt string, maxTokens int) (*Rsp, error) {
	message := Message{
		Role:    roleUser,
		Content: prompt,
	}

	requestBody := Req{
		Model:     model3dot5Turbo,
		Messages:  []Message{message},
		MaxTokens: maxTokens,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	// Retrieve the API key from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("API key not set in .env file")
	}

	url := "https://api.openai.com/v1/chat/completions"

	// Create and send the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to ChatGPT API: %w", err)
	}
	defer resp.Body.Close()

	// Reading and handling the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response Rsp
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	// Check if the API response contains an error
	if response.Error != nil {
		return nil, fmt.Errorf("API Error: %s - %s", response.Error.Code, response.Error.Message)

	}

	fmt.Printf("📃Raw Response:\n%s\n", string(body))

	return &response, nil
}
