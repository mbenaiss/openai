package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const url = "https://api.openai.com/v1/completions"

// Payload is the payload for the OpenAI API.
type Payload struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      float64  `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             float64  `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type response struct {
	Choices []choice `json:"choices"`
}

type choice struct {
	Text string `json:"text"`
}

// OpenAI is the OpenAI client.
type OpenAI struct {
	client *http.Client
	apiKey string
}

// New creates a new OpenAI client.
func New(apiKey string) *OpenAI {
	return &OpenAI{
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
		apiKey: apiKey,
	}
}

// Request is a wrapper around the OpenAI API.
func (o *OpenAI) Request(ctx context.Context, payload Payload) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("unable to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return "", fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+o.apiKey)

	res, err := o.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("unable to send request: %w", err)
	}

	defer res.Body.Close()

	var m response

	err = json.NewDecoder(res.Body).Decode(&m)
	if err != nil {
		return "", fmt.Errorf("unable to decode response: %w", err)
	}

	if len(m.Choices) == 0 {
		return "no results", nil
	}

	text := m.Choices[0].Text

	return text, nil
}
