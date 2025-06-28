package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	aiTypes "github.com/MRGHOSJ/docupocus/internal/ai/types"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	docTypes "github.com/MRGHOSJ/docupocus/internal/generator/types"
	docUtils "github.com/MRGHOSJ/docupocus/internal/generator/utils"
)

func GeneratePackageDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig) error {
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

	if len(existingContent) > 0 {
		b.Write(existingContent)
		b.WriteString("\n---\n\n")
	} else {
		// Package header with breadcrumbs
		b.WriteString(fmt.Sprintf("# 📦 Package: `%s`\n\n", pkg.Name))
		b.WriteString("[← Back to Overview](../README.md)\n\n")
	}

	// Add file-specific section
	b.WriteString(fmt.Sprintf("## 📄 File: `%s`\n\n", filepath.Base(filePath)))
	b.WriteString(fmt.Sprintf("> 📍 `%s`\n\n", docUtils.GetDisplayPath(filePath)))

	// TOC for package
	b.WriteString("## 📑 Contents\n\n")
	if len(pkg.Structs) > 0 {
		b.WriteString(fmt.Sprintf("- [🧱 Structs (%d)](#-structs)\n", len(pkg.Structs)))
	}
	if len(pkg.Funcs) > 0 {
		b.WriteString(fmt.Sprintf("- [🔧 Functions (%d)](#-functions)\n", len(pkg.Funcs)))
	}
	b.WriteString("\n")

	// Structs section with cards
	if len(pkg.Structs) > 0 {
		b.WriteString("## 🧱 Structs\n\n")
		for _, s := range pkg.Structs {
			b.WriteString(fmt.Sprintf("### `%s`\n\n", s.Name))
			b.WriteString("```go\n" + FormatStruct(s) + "\n```\n\n")
			b.WriteString(formatDocumentation(s.Doc))
			b.WriteString("\n---\n\n")
		}
	}

	// Functions section with expandable details
	if len(pkg.Funcs) > 0 {
		b.WriteString("## 🔧 Functions\n\n")
		for _, f := range pkg.Funcs {
			b.WriteString("<details>\n")
			b.WriteString(fmt.Sprintf("<summary><b><code>%s(%s)</code></b></summary>\n\n",
				f.Name, formatParams(f.Parameters)))
			b.WriteString(formatDocumentation(f.Doc))
			b.WriteString("\n</details>\n\n")
		}
	}

	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}

// formatDocumentation formats a full Documentation struct to Markdown
func formatDocumentation(doc aiTypes.Documentation) string {
	if doc.Summary == "" {
		return "_No documentation available._\n"
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("**Summary:** %s\n\n", doc.Summary))

	if len(doc.Parameters) > 0 {
		b.WriteString("**Parameters:**\n")
		for _, p := range doc.Parameters {
			b.WriteString(fmt.Sprintf("- `%s` (%s): %s\n", p.Name, p.Type, p.Description))
		}
		b.WriteString("\n")
	}

	if doc.Returns != "" {
		b.WriteString(fmt.Sprintf("**Returns:** %s\n\n", doc.Returns))
	}

	if doc.TimeComplexity != "" || doc.SpaceComplexity != "" {
		b.WriteString("**Complexity:**\n")
		if doc.TimeComplexity != "" {
			b.WriteString(fmt.Sprintf("- Time: %s\n", doc.TimeComplexity))
		}
		if doc.SpaceComplexity != "" {
			b.WriteString(fmt.Sprintf("- Space: %s\n", doc.SpaceComplexity))
		}
		b.WriteString("\n")
	}

	if doc.UsageExample != "" {
		b.WriteString("**Example:**\n")
		b.WriteString(fmt.Sprintf("```go\n%s\n```\n\n", doc.UsageExample))
	}

	if len(doc.EdgeCases) > 0 {
		b.WriteString("**Edge Cases:**\n")
		for _, e := range doc.EdgeCases {
			b.WriteString(fmt.Sprintf("- %s\n", e))
		}
		b.WriteString("\n")
	}

	return b.String()
}

func formatParams(params []analyzer.Parameter) string {
	var parts []string
	for _, p := range params {
		parts = append(parts, fmt.Sprintf("%s %s", p.Name, p.Type))
	}
	return strings.Join(parts, ", ")
}

func FormatFunction(f analyzer.Function) string {
	var b strings.Builder
	if f.Receiver != "" {
		b.WriteString(fmt.Sprintf("func (%s) %s(", f.Receiver, f.Name))
	} else {
		b.WriteString(fmt.Sprintf("func %s(", f.Name))
	}
	b.WriteString(formatParams(f.Parameters))
	b.WriteString(")")
	if len(f.Results) > 0 {
		b.WriteString(" " + formatParams(f.Results))
	}
	return b.String()
}

func FormatStruct(s analyzer.Struct) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("type %s struct {\n", s.Name))
	for _, f := range s.Fields {
		b.WriteString(fmt.Sprintf("\t%s %s %s\n", f.Name, f.Type, f.Tag))
	}
	b.WriteString("}")
	return b.String()
}
