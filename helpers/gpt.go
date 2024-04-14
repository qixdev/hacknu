package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func SendGPT(systemPrompt, message, model string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return "", nil
	}

	url := "https://api.openai.com/v1/chat/completions"

	messages := []Message{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: message},
	}

	requestData := OpenAIRequest{
		Model:       model,
		Messages:    messages,
		Temperature: 0.7, // Adjust for more or less randomness
	}

	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenAIResponse

	if err := json.Unmarshal(responseBody, &response); err != nil {
		panic(err)
	}

	return response.Choices[0].Message.Content, nil
}
