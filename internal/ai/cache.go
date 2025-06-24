package ai

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Cache handles reading/writing documentation from a local file cache.
type Cache struct {
	dir string
}

func NewCache(dir string) *Cache {
	os.MkdirAll(dir, 0755)
	return &Cache{dir: dir}
}

type CacheKey struct {
	Hash     SemanticHash
	Language string
}

func (c *Cache) filename(key CacheKey) string {
	hashStr := hex.EncodeToString(key.Hash[:])
	lang := key.Language
	return filepath.Join(c.dir, fmt.Sprintf("%s_%s.json", hashStr, lang))
}

func (c *Cache) Get(key CacheKey) (Documentation, bool) {
	path := c.filename(key)
	if _, err := os.Stat(path); err != nil {
		return Documentation{}, false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Documentation{}, false
	}

	var doc Documentation
	if err := json.Unmarshal(data, &doc); err != nil {
		return Documentation{}, false
	}

	return doc, true
}

func (c *Cache) Set(key CacheKey, doc Documentation) error {
	path := c.filename(key)

	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
