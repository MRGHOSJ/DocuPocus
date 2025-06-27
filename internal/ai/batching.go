package ai

import (
	"context"
	"fmt"
	"sync"
	"time"

	docType "github.com/MRGHOSJ/docupocus/internal/ai/types"
)

func (c *Client) processBatchWithRetry(ctx context.Context, inputs, languages []string) ([]docType.Documentation, error) {
	var lastErr error

	for attempt := 0; attempt < c.config.MaxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(time.Duration(attempt*attempt) * time.Second):
				c.logger.Info("Retry attempt %d after backoff", attempt)
			}
		}

		results, err := c.processBatch(ctx, inputs, languages)
		if err == nil {
			return results, nil
		}

		lastErr = err
		c.logger.Warn("Attempt %d/%d failed: %v", attempt+1, c.config.MaxRetries, err)
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

func (c *Client) processYAMLBatchWithRetry(ctx context.Context, inputs, languages []string) ([]docType.YAMLDocumentation, error) {
	var lastErr error

	for attempt := 0; attempt < c.config.MaxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(time.Duration(attempt*attempt) * time.Second):
				c.logger.Info("Retry attempt %d after backoff", attempt)
			}
		}

		results, err := c.processYAMLBatch(ctx, inputs, languages)
		if err == nil {
			return results, nil
		}

		lastErr = err
		c.logger.Warn("Attempt %d/%d failed: %v", attempt+1, c.config.MaxRetries, err)
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

func (c *Client) processBatch(ctx context.Context, inputs, languages []string) ([]docType.Documentation, error) {
	// Calculate token counts and filter skippable inputs
	tokenCounts := make([]int, len(inputs))
	for i, input := range inputs {
		if cheapSkipFilter(input) {
			tokenCounts[i] = 0
			continue
		}
		tokenCounts[i] = CountTokens(input)
	}

	// Group inputs by token budget
	groups := groupByTokenCounts(inputs, tokenCounts, c.config.TokenBudget)

	// Prepare results structure
	results := make([]docType.Documentation, len(inputs))
	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, 1)

	// Process each group concurrently
	for _, group := range groups {
		wg.Add(1)
		go func(indices []int) {
			defer wg.Done()

			// Get batch documentation
			batchDocs, err := c.callBatchAPI(ctx, inputs, languages, indices)
			if err != nil {
				select {
				case errChan <- fmt.Errorf("batch failed on indices %v: %w", indices, err):
				default:
				}
				return
			}

			// Safely store results
			mu.Lock()
			defer mu.Unlock()
			for i, idx := range indices {
				if i < len(batchDocs) { // Ensure we don't exceed batchDocs length
					results[idx] = batchDocs[i]
				}
			}
		}(group)
	}

	wg.Wait()
	close(errChan)

	// Check for any errors
	if err := <-errChan; err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) processYAMLBatch(ctx context.Context, inputs, languages []string) ([]docType.YAMLDocumentation, error) {
	tokenCounts := make([]int, len(inputs))
	for i, input := range inputs {
		if cheapSkipFilter(input) {
			tokenCounts[i] = 0
			continue
		}
		tokenCounts[i] = CountTokens(input)
	}

	groups := groupByTokenCounts(inputs, tokenCounts, c.config.TokenBudget)
	results := make([]docType.YAMLDocumentation, len(inputs))
	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, 1)

	for _, group := range groups {
		wg.Add(1)
		go func(indices []int) {
			defer wg.Done()

			// Get batch documentation
			rawDocs, err := c.callBatchYamlAPI(ctx, inputs, languages, indices)
			if err != nil {
				select {
				case errChan <- fmt.Errorf("batch failed on indices %v: %w", indices, err):
				default:
				}
				return
			}

			// Unmarshal each item into YAMLDocumentation
			mu.Lock()
			defer mu.Unlock()
			for i, idx := range indices {
				if i < len(rawDocs) {
					results[idx] = rawDocs[i]
				}
			}
		}(group)
	}

	wg.Wait()
	close(errChan)

	if err := <-errChan; err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) callBatchAPI(ctx context.Context, inputs, languages []string, indices []int) ([]docType.Documentation, error) {
	prompts := make([]string, len(indices))
	for i, idx := range indices {
		prompts[i] = c.buildPrompt(inputs[idx], languages[idx])
	}

	combinedPrompt := c.buildBatchPromptCodeAssistant(prompts)

	response, err := c.backend.Call(ctx, combinedPrompt)
	if err != nil {
		return nil, err
	}

	c.logger.Info("Raw batch API response: %s", response)

	return c.parseBatchResponse(response, len(prompts))
}

func (c *Client) callBatchYamlAPI(ctx context.Context, inputs, languages []string, indices []int) ([]docType.YAMLDocumentation, error) {
	prompts := make([]string, len(indices))
	for i, idx := range indices {
		prompts[i] = c.buildPrompt(inputs[idx], languages[idx])
	}

	combinedPrompt := c.buildBatchPromptYamlDocumentation(prompts)

	response, err := c.backend.Call(ctx, combinedPrompt)
	if err != nil {
		return nil, err
	}

	c.logger.Info("Raw batch API response: %s", response)

	return c.parseYAMLBatchResponse(response, len(prompts))
}
