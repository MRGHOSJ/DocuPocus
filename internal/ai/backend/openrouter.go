package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type OpenRouterBackend struct {
	httpClient  *http.Client
	config      BackendConfig
	rateLimiter *rate.Limiter
}

func NewOpenRouterBackend(cfg BackendConfig) *OpenRouterBackend {
	if cfg.Endpoint == "" {
		cfg.Endpoint = "https://openrouter.ai/api/v1"
	}

	if cfg.RateLimit <= 0 {
		cfg.RateLimit = 18
	}

	return &OpenRouterBackend{
		httpClient:  &http.Client{Timeout: 120 * time.Second},
		config:      cfg,
		rateLimiter: rate.NewLimiter(rate.Every(time.Minute/time.Duration(cfg.RateLimit)), cfg.RateLimit),
	}
}

func (b *OpenRouterBackend) Name() string {
	return "openrouter"
}

func (b *OpenRouterBackend) Call(ctx context.Context, prompt string) (string, error) {
	if err := b.rateLimiter.Wait(ctx); err != nil {
		return "", err
	}

	messages := []map[string]interface{}{
		{"role": "user", "content": prompt},
	}

	body := map[string]interface{}{
		"model":       b.config.Model,
		"messages":    messages,
		"max_tokens":  800,
		"temperature": 0.2,
	}

	resp, err := b.callAPI(ctx, body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var apiResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(apiResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return apiResponse.Choices[0].Message.Content, nil
}

func (b *OpenRouterBackend) callAPI(ctx context.Context, body interface{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", b.config.Endpoint+"/chat/completions", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+b.config.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "https://github.com/MRGHOSJ/docupocus")
	req.Header.Set("X-Title", "DocuPocus")

	return b.httpClient.Do(req)
}
