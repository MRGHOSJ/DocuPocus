package ai

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
)

// SemanticHash and CacheKey as before
type SemanticHash [32]byte

type CacheKey struct {
	Hash     SemanticHash
	Language string
}

type Cache struct {
	dir string
}

func NewCache(dir string) *Cache {
	os.MkdirAll(dir, 0755)
	return &Cache{dir: dir}
}

func GenerateSemanticHash(input string) SemanticHash {
	normalized := strings.Join(strings.Fields(strings.ToLower(input)), " ")
	return sha256.Sum256([]byte(normalized))
}

func (c *Cache) filename(key CacheKey) string {
	hashStr := hex.EncodeToString(key.Hash[:])
	lang := key.Language
	return filepath.Join(c.dir, hashStr+"_"+lang+".json")
}

// Get returns raw cached data bytes, or false if not found or error
func (c *Cache) Get(key CacheKey) ([]byte, bool) {
	path := c.filename(key)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	return data, true
}

// Set writes raw data bytes to cache
func (c *Cache) Set(key CacheKey, data []byte) error {
	path := c.filename(key)
	return os.WriteFile(path, data, 0644)
}

func GetDoc[T any](c *Cache, key CacheKey, unmarshal func([]byte, *T) error) (T, bool) {
	var zero T
	data, ok := c.Get(key)
	if !ok {
		return zero, false
	}
	var doc T
	if err := unmarshal(data, &doc); err != nil {
		return zero, false
	}
	return doc, true
}

func SetDoc[T any](c *Cache, key CacheKey, doc T, marshal func(T) ([]byte, error)) error {
	data, err := marshal(doc)
	if err != nil {
		return err
	}
	return c.Set(key, data)
}
