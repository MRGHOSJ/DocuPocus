package ai

import (
	"context"
	"encoding/json"
	"fmt"

	ai "github.com/MRGHOSJ/docupocus/internal/ai/backend"
	aiCache "github.com/MRGHOSJ/docupocus/internal/ai/cache"
	docType "github.com/MRGHOSJ/docupocus/internal/ai/types"
)

type Client struct {
	backend ai.Backend
	cache   *aiCache.Cache
	logger  Logger
	config  ai.BackendConfig
}

func NewClient(backend ai.Backend, cfg ai.BackendConfig) *Client {
	return &Client{
		backend: backend,
		cache:   aiCache.NewCache("ai-cache"),
		logger:  NewStdLogger(),
		config:  cfg,
	}
}

func (c *Client) ApplyDefaults() {
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
}

func (c *Client) EnhanceDocumentationBatch(ctx context.Context, inputs, languages []string) ([]docType.Documentation, error) {
	get := func(key aiCache.CacheKey) (docType.Documentation, bool) {
		return aiCache.GetDoc[docType.Documentation](c.cache, key, jsonUnmarshalAdapter[docType.Documentation])
	}
	set := func(key aiCache.CacheKey, doc docType.Documentation) error {
		return aiCache.SetDoc(c.cache, key, doc, jsonMarshalIndentAdapter[docType.Documentation])
	}

	return EnhanceGenericBatch(
		ctx, c, inputs, languages,
		get, set, c.processBatchWithRetry,
	)
}

func (c *Client) EnhanceYAMLDocumentationBatch(ctx context.Context, inputs, languages []string) ([]docType.YAMLDocumentation, error) {
	get := func(key aiCache.CacheKey) (docType.YAMLDocumentation, bool) {
		return aiCache.GetDoc[docType.YAMLDocumentation](c.cache, key, jsonUnmarshalAdapter[docType.YAMLDocumentation])
	}
	set := func(key aiCache.CacheKey, doc docType.YAMLDocumentation) error {
		return aiCache.SetDoc(c.cache, key, doc, jsonMarshalIndentAdapter[docType.YAMLDocumentation])
	}

	return EnhanceGenericBatch(
		ctx, c, inputs, languages,
		get, set, c.processYAMLBatchWithRetry,
	)
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

func (c *Client) deduplicateInputs(inputs, languages []string) (
	hashes []aiCache.SemanticHash,
	uniqueInputs []string,
	uniqueLanguages []string,
	reverseMap []int,
) {
	type inputKey struct {
		hash     aiCache.SemanticHash
		language string
	}

	hashes = make([]aiCache.SemanticHash, len(inputs))
	uniqueMap := make(map[inputKey]int)
	reverseMap = make([]int, len(inputs))

	for i, input := range inputs {
		hash := aiCache.GenerateSemanticHash(input)
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

	return
}

func EnhanceGenericBatch[T any](
	ctx context.Context,
	c *Client,
	inputs, languages []string,
	getCache func(aiCache.CacheKey) (T, bool),
	setCache func(aiCache.CacheKey, T) error,
	callBatch func(context.Context, []string, []string) ([]T, error),
) ([]T, error) {
	if len(inputs) == 0 {
		c.logger.Info("No inputs provided.")
		return []T{}, nil
	}

	if len(inputs) != len(languages) {
		return nil, fmt.Errorf("inputs and languages length mismatch: %d vs %d", len(inputs), len(languages))
	}

	c.logger.Info("🚀 Starting batch documentation enhancement for %d inputs", len(inputs))

	hashes, uniqueInputs, uniqueLanguages, reverseMap := c.deduplicateInputs(inputs, languages)
	c.logger.Info("📦 %d unique inputs after deduplication", len(uniqueInputs))

	cachedResults := make([]T, len(uniqueInputs))
	toProcess := []int{}

	for i := range uniqueInputs {
		key := aiCache.CacheKey{
			Hash:     hashes[i],
			Language: uniqueLanguages[i],
		}
		if cached, ok := getCache(key); ok {
			c.logger.Info("✅ Cache hit for input %d", i)
			cachedResults[i] = cached
		} else {
			c.logger.Info("❌ Cache miss for input %d", i)
			toProcess = append(toProcess, i)
		}
	}

	if len(toProcess) > 0 {
		batchSize := c.config.BatchSize
		for start := 0; start < len(toProcess); start += batchSize {
			end := min(start+batchSize, len(toProcess))
			indices := toProcess[start:end]

			c.logger.Info("🔄 Processing batch %d–%d", start, end-1)
			processInputs := make([]string, len(indices))
			processLangs := make([]string, len(indices))
			for i, idx := range indices {
				processInputs[i] = uniqueInputs[idx]
				processLangs[i] = uniqueLanguages[idx]
			}

			batchDocs, err := callBatch(ctx, processInputs, processLangs)
			if err != nil {
				return nil, fmt.Errorf("batch %d–%d failed: %w", start, end, err)
			}

			for i, idx := range indices {
				cachedResults[idx] = batchDocs[i]
				_ = setCache(aiCache.CacheKey{
					Hash:     hashes[idx],
					Language: uniqueLanguages[idx],
				}, batchDocs[i])
			}
		}
	}

	// Restore original order
	finalResults := make([]T, len(inputs))
	for i, idx := range reverseMap {
		if idx < len(cachedResults) {
			finalResults[i] = cachedResults[idx]
		}
	}

	return finalResults, nil
}

func jsonUnmarshalAdapter[T any](data []byte, v *T) error {
	return json.Unmarshal(data, v)
}

func jsonMarshalIndentAdapter[T any](v T) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
