package ai

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// ExtractJSONArray finds the first JSON array substring in raw text.
func ExtractJSONArray(raw string) (string, error) {
	re := regexp.MustCompile(`\[[\s\S]*\]`)
	match := re.FindString(raw)
	if match == "" {
		return "", fmt.Errorf("no JSON array found in response")
	}
	return match, nil
}

// ParseObjectKeysAsArray extracts map keys as string slice as a fallback parser.
func ParseObjectKeysAsArray(raw string, expectedCount int) ([]string, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &obj); err != nil {
		return nil, fmt.Errorf("fallback: failed to parse JSON object: %w", err)
	}

	expls := make([]string, 0, expectedCount)

	re := regexp.MustCompile(`explanation(\d+)`)
	type kv struct {
		key   string
		index int
	}
	var keys []kv
	for k := range obj {
		if m := re.FindStringSubmatch(k); len(m) == 2 {
			var idx int
			fmt.Sscanf(m[1], "%d", &idx)
			keys = append(keys, kv{key: k, index: idx})
		} else {
			keys = append(keys, kv{key: k, index: 0})
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].index < keys[j].index
	})

	for _, k := range keys {
		v := obj[k.key]
		if s, ok := v.(string); ok && s != "" {
			expls = append(expls, s)
		} else if s, ok := v.(interface{}); ok {
			if str, ok := s.(string); ok && str != "" {
				expls = append(expls, str)
			}
		}
	}

	if len(expls) == 0 {
		for _, k := range keys {
			expls = append(expls, k.key)
		}
	}

	if len(expls) != expectedCount {
		return nil, fmt.Errorf("fallback explanation count mismatch: expected %d, got %d", expectedCount, len(expls))
	}
	return expls, nil
}

// MapResults maps cached and new results to the original input order.
func MapResults(cachedResults []string, reverseMap []int, inputLen int) []string {
	finalResults := make([]string, inputLen)
	for i, idx := range reverseMap {
		finalResults[i] = cachedResults[idx]
	}
	return finalResults
}

// EstimateTokens provides a rough estimate of token count for an input string.
func EstimateTokens(input string) int {
	words := len(strings.Fields(input))
	return int(float64(words) * 1.5) // Rough estimate: 1 word ≈ 1.5 tokens
}
