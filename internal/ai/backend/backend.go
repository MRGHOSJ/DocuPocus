package ai

import "context"

type Backend interface {
	Call(ctx context.Context, prompt string) (string, error)

	Name() string
}

type BackendConfig struct {
	Model       string
	Endpoint    string
	APIKey      string
	RateLimit   int
	MaxRetries  int
	BatchSize   int
	TokenBudget int
	Prompt      string
}
