package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/analyzer"
)

func GetDisplayPath(fullPath string) string {
	parts := strings.Split(filepath.Clean(fullPath), string(filepath.Separator))
	if len(parts) < 2 {
		return fullPath
	}
	return filepath.Join(parts[len(parts)-2:]...)
}

func GetLanguage(path string) string {
	switch {
	case strings.HasSuffix(path, ".go"):
		return "Go"
	case strings.HasSuffix(path, ".py"):
		return "Python"
	case strings.HasSuffix(path, ".js"), strings.HasSuffix(path, ".ts"):
		return "JavaScript"
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		return "YAML"
	default:
		return "Unknown"
	}
}

func NormalizeFields(fields []analyzer.Field) []analyzer.Field {
	var normalized []analyzer.Field

	for _, f := range fields {
		if f.Type == "array" {
			itemsMap := make(map[int]analyzer.Field)
			maxIdx := -1

			for _, item := range f.Fields {
				var idx int
				// Match index-based names like "[0]", "[1]", etc.
				if _, err := fmt.Sscanf(item.Name, "[%d]", &idx); err == nil {
					// Clear name to avoid printing as map key in YAML example
					item.Name = ""
					// Recursively normalize nested fields inside array item
					item.Fields = NormalizeFields(item.Fields)
					itemsMap[idx] = item
					if idx > maxIdx {
						maxIdx = idx
					}
				} else {
					// If item.Name doesn't match index pattern, just keep it as is,
					// recursively normalize children
					item.Fields = NormalizeFields(item.Fields)
					// Append directly to normalized fields (fallback)
					normalized = append(normalized, item)
				}
			}

			// Create ordered list from map by index
			orderedItems := make([]analyzer.Field, 0, maxIdx+1)
			for i := 0; i <= maxIdx; i++ {
				if item, exists := itemsMap[i]; exists {
					orderedItems = append(orderedItems, item)
				}
			}

			f.Fields = orderedItems
		} else if len(f.Fields) > 0 {
			// Recursively normalize children fields
			f.Fields = NormalizeFields(f.Fields)
		}

		normalized = append(normalized, f)
	}

	return normalized
}

// Helper to get value or placeholder string
func GetValueOrPlaceholder(f analyzer.Field) string {
	if f.Value != "" {
		return f.Value
	}
	switch f.Type {
	case "array":
		return "[]"
	case "map":
		return "{}"
	default:
		return "# TODO: Add value"
	}
}

func CalculateDocCompletion(pkg analyzer.Package) int {
	totalItems := len(pkg.Structs) + len(pkg.Funcs)
	if totalItems == 0 {
		return 100
	}

	documentedItems := 0
	for _, s := range pkg.Structs {
		if s.Doc.Summary != "" {
			documentedItems++
		}
	}
	for _, f := range pkg.Funcs {
		if f.Doc.Summary != "" {
			documentedItems++
		}
	}

	return (documentedItems * 100) / totalItems
}
