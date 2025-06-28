package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	docTypes "github.com/MRGHOSJ/docupocus/internal/generator/types"
	docUtils "github.com/MRGHOSJ/docupocus/internal/generator/utils"
)

func GenerateYAMLDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig) error {
	pkgDir := filepath.Join(cfg.OutputDir, pkg.Name)
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return fmt.Errorf("failed to create package directory: %w", err)
	}

	readmePath := filepath.Join(pkgDir, "README.md")

	var existingContent []byte
	if _, err := os.Stat(readmePath); err == nil {
		existingContent, _ = os.ReadFile(readmePath)
	}

	var b strings.Builder

	// Header
	if len(existingContent) > 0 {
		b.Write(existingContent)
		b.WriteString("\n---\n\n")
	} else {
		b.WriteString(fmt.Sprintf("# 📄 YAML Configuration: `%s`\n\n", pkg.Name))
		b.WriteString("[← Back to Overview](../README.md)\n\n")
	}

	b.WriteString(fmt.Sprintf("## 📄 File: `%s`\n\n", filepath.Base(filePath)))
	b.WriteString(fmt.Sprintf("> 📍 Path: `%s`\n\n", docUtils.GetDisplayPath(filePath)))

	for _, s := range pkg.Structs {
		doc := s.DocYAML

		// Expandable Resource Summary
		if doc.Summary != "" {
			b.WriteString("🚀 Resource Summary\n\n")
			b.WriteString(fmt.Sprintf("- **Kind:** `%s`\n", s.Name))
			b.WriteString(fmt.Sprintf("- **Description:** %s\n\n", doc.Summary))
		}

		// Expandable Configuration Example
		if len(s.Fields) > 0 {
			normalizedFields := docUtils.NormalizeFields(s.Fields)
			b.WriteString("<details>\n")
			b.WriteString(fmt.Sprintf("<summary>⚙️ Configuration Example for `%s`</summary>\n\n", s.Name))
			b.WriteString("```yaml\n")
			b.WriteString(generateYAMLExample(normalizedFields, 0))
			b.WriteString("```\n")
			b.WriteString("</details>\n\n")
		}

		// Expandable Field Reference (each field in its own collapsible)
		if len(doc.Fields) > 0 {
			b.WriteString("<details>\n")
			b.WriteString("<summary>📑 Field Reference</summary>\n\n")
			for _, f := range doc.Fields {
				b.WriteString("<details>\n")
				b.WriteString(fmt.Sprintf("<summary>`%s`</summary>\n\n", f.Name))
				b.WriteString(fmt.Sprintf("- **Type:** `%s`\n", f.Type))
				if f.Description != "" {
					b.WriteString(fmt.Sprintf("- **Description:** %s\n", f.Description))
				}
				b.WriteString("</details>\n\n")
			}
			b.WriteString("</details>\n\n")
		}

		// Expandable Examples
		if len(doc.Examples) > 0 {
			b.WriteString("<details>\n")
			b.WriteString("<summary>🔍 Examples</summary>\n\n```yaml\n")
			for key, val := range doc.Examples {
				b.WriteString(fmt.Sprintf("%s: %v\n", key, val))
			}
			b.WriteString("```\n")
			b.WriteString("</details>\n\n")
		}

		// Expandable Defaults
		if len(doc.Defaults) > 0 {
			b.WriteString("<details>\n")
			b.WriteString("<summary>🌐 Defaults</summary>\n\n")
			for key, val := range doc.Defaults {
				b.WriteString(fmt.Sprintf("- **%s**: `%v`\n", key, val))
			}
			b.WriteString("</details>\n\n")
		}

		// Expandable Usage
		if doc.Usage != "" {
			b.WriteString("<details>\n")
			b.WriteString("<summary>🧰 Usage</summary>\n\n")
			b.WriteString(doc.Usage + "\n")
			b.WriteString("</details>\n\n")
		}

		// Expandable Edge Cases
		if len(doc.BestPractices) > 0 {
			b.WriteString("<details>\n")
			b.WriteString("<summary>⚠️ Edge Cases</summary>\n\n")
			for _, ec := range doc.BestPractices {
				b.WriteString(fmt.Sprintf("- %s\n", ec))
			}
			b.WriteString("</details>\n\n")
		}
	}

	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}

func generateYAMLExample(fields []analyzer.Field, indentLevel int) string {
	var b strings.Builder
	baseIndent := strings.Repeat("  ", indentLevel)

	for _, f := range fields {
		switch f.Type {
		case "array":
			// Write array key if present
			if f.Name != "" {
				b.WriteString(fmt.Sprintf("%s%s:\n", baseIndent, f.Name))
			}

			// If no items, add empty array placeholder
			if len(f.Fields) == 0 {
				b.WriteString(fmt.Sprintf("%s- # TODO: Add item\n", baseIndent+"  ")) // indent array item one more level
				continue
			}

			// Process array items
			for _, item := range f.Fields {
				// Start list item
				b.WriteString(fmt.Sprintf("%s- ", baseIndent+"  "))

				// Case 1: item has nested fields -> complex object
				if len(item.Fields) > 0 {
					if item.Name != "" {
						b.WriteString(fmt.Sprintf("%s:\n", item.Name))
					} else {
						b.WriteString("\n")
					}
					// recursively generate nested yaml with indentLevel+2 (one for itemIndent, one for nested)
					b.WriteString(generateYAMLExample(item.Fields, indentLevel+2))
					continue
				}

				// Case 2: simple key-value pair inside array item
				if item.Name != "" {
					b.WriteString(fmt.Sprintf("%s\n", fmt.Sprintf("%s: %s", item.Name, docUtils.GetValueOrPlaceholder(item))))
				} else {
					// Case 3: plain value in array item
					b.WriteString(fmt.Sprintf("%s\n", docUtils.GetValueOrPlaceholder(item)))
				}
			}

		case "map":
			if f.Name != "" {
				b.WriteString(fmt.Sprintf("%s%s:\n", baseIndent, f.Name))
			}
			if len(f.Fields) > 0 {
				b.WriteString(generateYAMLExample(f.Fields, indentLevel+1))
			} else {
				b.WriteString(fmt.Sprintf("%s  %s\n", baseIndent, docUtils.GetValueOrPlaceholder(f)))
			}

		default: // scalar value
			if f.Name != "" {
				b.WriteString(fmt.Sprintf("%s%s: %s\n", baseIndent, f.Name, docUtils.GetValueOrPlaceholder(f)))
			} else {
				b.WriteString(fmt.Sprintf("%s%s\n", baseIndent, docUtils.GetValueOrPlaceholder(f)))
			}
		}
	}

	return b.String()
}
