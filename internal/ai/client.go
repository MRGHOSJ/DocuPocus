package ai

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	ai "github.com/MRGHOSJ/docupocus/internal/ai/backend"
)

type Client struct {
	backend ai.Backend
	cache   *Cache
	logger  Logger
	config  ai.BackendConfig
}

// Documentation represents the comprehensive documentation structure
type Documentation struct {
	Summary         string   `json:"summary"`
	Parameters      []Param  `json:"parameters"`
	Returns         string   `json:"returns"`
	TimeComplexity  string   `json:"time_complexity"`
	SpaceComplexity string   `json:"space_complexity"`
	UsageExample    string   `json:"usage_example"`
	EdgeCases       []string `json:"edge_cases"`
}

// Param describes a function parameter
type Param struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type SemanticHash [32]byte

func NewClient(backend ai.Backend, cfg ai.BackendConfig) *Client {
	// Set default values
	if cfg.RateLimit <= 0 {
		cfg.RateLimit = 18
	}
	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = 3
	}
	if cfg.BatchSize <= 0 {
		cfg.BatchSize = 3
	}
	if cfg.TokenBudget <= 0 {
		cfg.TokenBudget = 7000
	}

	return &Client{
		backend: backend,
		cache:   NewCache("ai-cache"),
		logger:  NewStdLogger(),
		config:  cfg,
	}
}

func (c *Client) EnhanceDocumentationBatch(ctx context.Context, inputs, languages []string) ([]Documentation, error) {
	if len(inputs) == 0 {
		c.logger.Info("No inputs provided for documentation enhancement.")
		return []Documentation{}, nil
	}

	if len(inputs) != len(languages) {
		return nil, fmt.Errorf("inputs and languages length mismatch: %d vs %d", len(inputs), len(languages))
	}

	c.logger.Info("🚀 Starting batch documentation enhancement for %d inputs", len(inputs))

	// Apply defaults (defensive programming)
	if c.config.BatchSize <= 0 {
		c.config.BatchSize = 3
	}
	if c.config.TokenBudget <= 0 {
		c.config.TokenBudget = 7000
	}
	if c.config.RateLimit <= 0 {
		c.config.RateLimit = 18
	}
	if c.config.MaxRetries <= 0 {
		c.config.MaxRetries = 3
	}

	// Generate semantic hashes and map unique inputs
	c.logger.Info("🔍 Deduplicating inputs and generating semantic hashes...")
	type inputKey struct {
		hash     SemanticHash
		language string
	}

	hashes := make([]SemanticHash, len(inputs))
	uniqueMap := make(map[inputKey]int)
	reverseMap := make([]int, len(inputs))
	uniqueInputs := make([]string, 0)
	uniqueLanguages := make([]string, 0)

	for i, input := range inputs {
		hash := generateSemanticHash(input)
		hashes[i] = hash

		key := inputKey{hash, languages[i]}
		if idx, exists := uniqueMap[key]; exists {
			reverseMap[i] = idx
			continue
		}
		uniqueMap[key] = len(uniqueInputs)
		reverseMap[i] = len(uniqueInputs)
		uniqueInputs = append(uniqueInputs, input)
		uniqueLanguages = append(uniqueLanguages, languages[i])
	}

	c.logger.Info("📦 %d unique inputs after deduplication", len(uniqueInputs))

	// Check the cache
	c.logger.Info("📁 Checking cache for existing documentation...")
	cachedResults := make([]Documentation, len(uniqueInputs))
	toProcess := make([]int, 0)

	for i := range uniqueInputs {
		cacheKey := CacheKey{
			Hash:     hashes[i],
			Language: uniqueLanguages[i],
		}
		if cached, ok := c.cache.Get(cacheKey); ok {
			c.logger.Info("✅ Cache hit for input %d (hash: %x)", i, hashes[i])
			cachedResults[i] = cached
		} else {
			c.logger.Error("❌ Cache miss for input %d (hash: %x)", i, hashes[i])
			toProcess = append(toProcess, i)
		}
	}

	c.logger.Info("📤 %d inputs to be processed by AI", len(toProcess))

	// Process uncached in batches
	if len(toProcess) > 0 {
		batchSize := c.config.BatchSize
		for start := 0; start < len(toProcess); start += batchSize {
			end := start + batchSize
			if end > len(toProcess) {
				end = len(toProcess)
			}
			indices := toProcess[start:end]

			c.logger.Info("🔄 Processing batch %d–%d (size %d)", start, end-1, len(indices))

			// Extract batch inputs
			processInputs := make([]string, len(indices))
			processLangs := make([]string, len(indices))
			for i, idx := range indices {
				processInputs[i] = uniqueInputs[idx]
				processLangs[i] = uniqueLanguages[idx]
			}

			// Retryable batch processing
			docs, err := c.processBatchWithRetry(ctx, processInputs, processLangs)
			if err != nil {
				c.logger.Error("❌ Batch processing failed (indices %v): %v", indices, err)
				return nil, fmt.Errorf("batch %d–%d failed: %w", start, end, err)
			}

			// Store in cache and results
			for i, idx := range indices {
				cachedResults[idx] = docs[i]

				cacheKey := CacheKey{
					Hash:     hashes[idx],
					Language: uniqueLanguages[idx],
				}
				if err := c.cache.Set(cacheKey, docs[i]); err != nil {
					c.logger.Warn("⚠️ Failed to cache result for input %d: %v", idx, err)
				}
			}
		}
	}

	// Reconstruct results in original input order
	finalResults := make([]Documentation, len(inputs))
	for i, idx := range reverseMap {
		if idx < len(cachedResults) {
			finalResults[i] = cachedResults[idx]
		} else {
			c.logger.Warn("⚠️ Missing cached result for index %d (mapped to %d)", i, idx)
			finalResults[i] = Documentation{Summary: "Documentation generation failed"}
		}
	}

	c.logger.Info("✅ Enhancement complete: %d total, %d unique, %d uncached", len(inputs), len(uniqueInputs), len(toProcess))
	return finalResults, nil
}

func (c *Client) processBatchWithRetry(ctx context.Context, inputs, languages []string) ([]Documentation, error) {
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

func (c *Client) processBatch(ctx context.Context, inputs, languages []string) ([]Documentation, error) {
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
	results := make([]Documentation, len(inputs))
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

func (c *Client) callBatchAPI(ctx context.Context, inputs, languages []string, indices []int) ([]Documentation, error) {
	prompts := make([]string, len(indices))
	for i, idx := range indices {
		prompts[i] = c.buildPrompt(inputs[idx], languages[idx])
	}

	combinedPrompt := c.buildBatchPrompt(prompts)

	response, err := c.backend.Call(ctx, combinedPrompt)
	if err != nil {
		return nil, err
	}

	c.logger.Info("Raw batch API response: %s", response)

	return c.parseBatchResponse(response, len(prompts))
}

func (c *Client) buildPrompt(input, language string) string {
	if c.config.Prompt == "" {
		return fmt.Sprintf("Generate a concise explanation (10-15 words) for this %s code:\n```%s\n%s\n```",
			language, language, input)
	}
	return fmt.Sprintf(c.config.Prompt, language, language, input)
}
func (c *Client) buildBatchPrompt(prompts []string) string {
	var sb strings.Builder

	// Enhanced prompt with multiple documentation requirements
	sb.WriteString("You are an **advanced code documentation assistant**. For each code snippet provided, generate:\n\n")
	sb.WriteString("- A concise functional summary (15-20 words)\n")
	sb.WriteString("- Key parameters/inputs with types\n")
	sb.WriteString("- Return value/output description\n")
	sb.WriteString("- Time complexity analysis (Big-O notation)\n")
	sb.WriteString("- Space complexity analysis\n")
	sb.WriteString("- One common usage example\n")
	sb.WriteString("- Potential edge cases to consider\n\n")

	sb.WriteString(`**Return a JSON array** where each element is **an object containing these documentation aspects**.`)
	sb.WriteString("\n\nCode snippets:\n[\n")

	for i, prompt := range prompts {
		// Properly escape JSON and trim whitespace
		cleanPrompt := strings.TrimSpace(prompt)
		escapedPrompt := strings.ReplaceAll(cleanPrompt, `"`, `\"`)
		sb.WriteString(fmt.Sprintf(`  "Snippet %d: %s"`, i+1, escapedPrompt))
		if i < len(prompts)-1 {
			sb.WriteString(",\n")
		}
	}

	sb.WriteString("\n]\n\n")
	sb.WriteString(`**Return format example**:
[
  {
    "summary": "Function that adds two integers",
    "parameters": [
      {"name": "a", "type": "int", "description": "First operand"},
      {"name": "b", "type": "int", "description": "Second operand"}
    ],
    "returns": "Sum of the two integers as int",
    "time_complexity": "O(1)",
    "space_complexity": "O(1)",
    "usage_example": "sum := add(3, 5) // returns 8",
    "edge_cases": [
      "Integer overflow with large numbers",
      "Undefined behavior with extremely large values"
    ]
  }
]`)

	return sb.String()
}

func (c *Client) parseBatchResponse(response string, expectedCount int) ([]Documentation, error) {
	// First try to parse as proper JSON array
	var docs []Documentation
	if err := json.Unmarshal([]byte(response), &docs); err == nil {
		if len(docs) == expectedCount {
			return docs, nil
		}
		return nil, fmt.Errorf("documentation count mismatch: expected %d, got %d", expectedCount, len(docs))
	}

	// Fallback to extracting JSON from text response
	cleaned, err := ExtractJSONArray(response)
	if err != nil {
		return nil, fmt.Errorf("failed to extract JSON from response: %v", err)
	}

	if err := json.Unmarshal([]byte(cleaned), &docs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal documentation: %v", err)
	}

	if len(docs) != expectedCount {
		return nil, fmt.Errorf("documentation count mismatch: expected %d, got %d", expectedCount, len(docs))
	}

	return docs, nil
}

func generateSemanticHash(input string) SemanticHash {
	normalized := strings.Join(strings.Fields(strings.ToLower(input)), " ")
	return sha256.Sum256([]byte(normalized))
}

func groupByTokenCounts(inputs []string, tokenCounts []int, budget int) [][]int {
	var groups [][]int
	var currentGroup []int
	currentTokens := 0

	for i, count := range tokenCounts {
		if cheapSkipFilter(inputs[i]) {
			continue
		}

		if currentTokens+count > budget && len(currentGroup) > 0 {
			groups = append(groups, currentGroup)
			currentGroup = nil
			currentTokens = 0
		}

		currentGroup = append(currentGroup, i)
		currentTokens += count
	}

	if len(currentGroup) > 0 {
		groups = append(groups, currentGroup)
	}

	return groups
}
