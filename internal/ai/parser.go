package ai

import (
	"encoding/json"
	"fmt"

	docType "github.com/MRGHOSJ/docupocus/internal/ai/types"
)

func (c *Client) parseBatchResponse(response string, expectedCount int) ([]docType.Documentation, error) {
	// First try to parse as proper JSON array
	var docs []docType.Documentation
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

func (c *Client) parseYAMLBatchResponse(response string, expectedCount int) ([]docType.YAMLDocumentation, error) {
	var docs []docType.YAMLDocumentation

	// Attempt to unmarshal directly
	if err := json.Unmarshal([]byte(response), &docs); err == nil {
		if len(docs) == expectedCount {
			return docs, nil
		}
		return nil, fmt.Errorf("documentation count mismatch: expected %d, got %d", expectedCount, len(docs))
	}

	// Try extracting valid JSON array from text (fallback)
	cleaned, err := ExtractJSONArray(response)
	if err != nil {
		return nil, fmt.Errorf("failed to extract JSON array from response: %v", err)
	}

	if err := json.Unmarshal([]byte(cleaned), &docs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML documentation: %v", err)
	}

	if len(docs) != expectedCount {
		return nil, fmt.Errorf("documentation count mismatch: expected %d, got %d", expectedCount, len(docs))
	}

	return docs, nil
}
