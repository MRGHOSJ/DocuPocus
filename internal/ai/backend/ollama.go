package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/ollama/ollama/api"
)

type OllamaBackend struct {
	client *api.Client
	config BackendConfig
}

func NewOllamaBackend(cfg BackendConfig) (*OllamaBackend, error) {
	cli, err := api.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	return &OllamaBackend{
		client: cli,
		config: cfg,
	}, nil
}

func (b *OllamaBackend) Name() string {
	return "ollama"
}

func (b *OllamaBackend) Call(ctx context.Context, prompt string) (string, error) {
	var response strings.Builder

	err := b.client.Generate(ctx, &api.GenerateRequest{
		Model:  b.config.Model,
		Prompt: prompt,
	}, func(gr api.GenerateResponse) error {
		response.WriteString(gr.Response)
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("generation failed: %w", err)
	}
	return response.String(), nil
}
