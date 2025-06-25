package generator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"
)

type GeneratorConfig struct {
	AIClient  *ai.Client
	OutputDir string
	Project   struct {
		Name        string
		Description string
		RepoURL     string
	}
}

type AIRequest struct {
	Input    string
	Language string
	Target   *ai.Documentation
}

func GeneratePackageDocs(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig) error {
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create doc directory: %w", err)
	}

	// Generate project-wide documentation
	if err := generateProjectReadme(result, cfg); err != nil {
		return fmt.Errorf("failed to generate project README: %w", err)
	}

	// Generate sidebar for navigation
	if err := generateSidebar(result, cfg); err != nil {
		return fmt.Errorf("failed to generate sidebar: %w", err)
	}

	// Process AI enhancements
	var aiRequests []AIRequest
	for fi := range result.Files {
		for pi := range result.Files[fi].Packages {
			pkg := &result.Files[fi].Packages[pi]
			lang := getLanguage(result.Files[fi].Path)

			for si := range pkg.Structs {
				if cfg.AIClient != nil {
					input := formatStruct(pkg.Structs[si])
					pkg.Structs[si].Doc = ai.Documentation{}
					aiRequests = append(aiRequests, AIRequest{
						Input:    input,
						Language: lang,
						Target:   &pkg.Structs[si].Doc,
					})
				}
			}
			for fi := range pkg.Funcs {
				if cfg.AIClient != nil {
					input := formatFunction(pkg.Funcs[fi])
					pkg.Funcs[fi].Doc = ai.Documentation{}
					aiRequests = append(aiRequests, AIRequest{
						Input:    input,
						Language: lang,
						Target:   &pkg.Funcs[fi].Doc,
					})
				}
			}
		}
	}

	if cfg.AIClient != nil && len(aiRequests) > 0 {
		fmt.Printf("🚀 Sending %d documentation enhancements to AI\n", len(aiRequests))
		processAIRequests(aiRequests, cfg.AIClient)
	}

	// Generate package documentation
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			if err := generatePackageDoc(pkg, file.Path, cfg); err != nil {
				return fmt.Errorf("failed to generate docs for package %s: %w", pkg.Name, err)
			}
		}
	}

	return nil
}

func generateProjectReadme(result *analyzer.AnalyzerResult, cfg GeneratorConfig) error {
	var b strings.Builder
	readmePath := filepath.Join(cfg.OutputDir, "README.md")

	// Project header with emojis and badges
	b.WriteString(fmt.Sprintf("# 📘 %s\n\n", cfg.Project.Name))
	b.WriteString(fmt.Sprintf("> %s\n\n", cfg.Project.Description))
	b.WriteString(fmt.Sprintf("[![Go](https://img.shields.io/badge/Go-%%20%%E2%%9D%%A4%%EF%%B8%%8F-blue)](%s) ", cfg.Project.RepoURL))
	b.WriteString(fmt.Sprintf("[![GitHub](https://img.shields.io/badge/GitHub-Repository-lightgrey)](%s)\n\n", cfg.Project.RepoURL))

	// TOC with emojis
	b.WriteString("## 📚 Table of Contents\n\n")
	b.WriteString("- [✨ Project Overview](#-project-overview)\n")
	b.WriteString("- [📦 Packages](#-packages)\n")
	b.WriteString("- [🚀 Quick Start](#-quick-start)\n")
	b.WriteString("- [💡 Best Practices](#-best-practices)\n\n")

	// Overview section with columns using tables
	b.WriteString("## ✨ Project Overview\n\n")
	b.WriteString("<table>\n<tr>\n<td valign=\"top\" width=\"50%%\">\n\n")
	b.WriteString("### 🎯 Key Features\n\n")
	b.WriteString("- Feature 1\n- Feature 2\n- Feature 3\n\n")
	b.WriteString("</td>\n<td valign=\"top\" width=\"50%%\">\n\n")
	b.WriteString("### 🛠️ Tech Stack\n\n")
	b.WriteString("- Go\n- Docker\n- gRPC\n\n")
	b.WriteString("</td>\n</tr>\n</table>\n\n")

	// Packages section with cards using tables
	b.WriteString("## 📦 Packages\n\n")
	b.WriteString("<table>\n")
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			docPath := filepath.Join(pkg.Name, "README.md")
			b.WriteString("<tr>\n<td valign=\"top\" width=\"33%%\">\n\n")
			b.WriteString(fmt.Sprintf("### [%s](%s)\n\n", pkg.Name, docPath))
			b.WriteString(fmt.Sprintf("`%s`\n\n", file.Path))
			b.WriteString(fmt.Sprintf("- %d structs\n", len(pkg.Structs)))
			b.WriteString(fmt.Sprintf("- %d functions\n", len(pkg.Funcs)))
			b.WriteString(fmt.Sprintf("- 📊 %d%% documented\n", calculateDocCompletion(pkg)))
			b.WriteString("\n</td>\n")
		}
	}
	b.WriteString("</tr>\n</table>\n\n")

	// Quick Start with tabs using details/summary
	b.WriteString("## 🚀 Quick Start\n\n")
	b.WriteString("<details>\n<summary><b>Local Development</b></summary>\n\n")
	b.WriteString("```bash\ngit clone " + cfg.Project.RepoURL + "\ncd project\ngo run main.go\n```\n\n")
	b.WriteString("</details>\n\n")
	b.WriteString("<details>\n<summary><b>Docker</b></summary>\n\n")
	b.WriteString("```bash\ndocker build -t myapp .\ndocker run -p 8080:8080 myapp\n```\n\n")
	b.WriteString("</details>\n\n")
	b.WriteString("<details>\n<summary><b>Cloud Deployment</b></summary>\n\n")
	b.WriteString("```bash\ngcloud app deploy\ngcloud app browse\n```\n\n")
	b.WriteString("</details>\n\n")

	// Best Practices
	b.WriteString("## 💡 Best Practices\n\n")
	b.WriteString("```diff\n+ Do\n- Don't\n```\n\n")
	b.WriteString("- ✅ Keep functions small and focused\n")
	b.WriteString("- ✅ Write clear documentation\n")
	b.WriteString("- ❌ Avoid global variables\n")
	b.WriteString("- ❌ Don't ignore errors\n\n")

	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}

func generatePackageDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig) error {
	var b strings.Builder
	pkgDir := filepath.Join(cfg.OutputDir, pkg.Name)

	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return fmt.Errorf("failed to create package directory: %w", err)
	}

	// Package header with breadcrumbs
	b.WriteString(fmt.Sprintf("# 📦 Package: `%s`\n\n", pkg.Name))
	b.WriteString(fmt.Sprintf("> 📍 `%s`\n\n", filePath))
	b.WriteString("[← Back to Overview](../README.md)\n\n")

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
			b.WriteString("```go\n" + formatStruct(s) + "\n```\n\n")
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

	readmePath := filepath.Join(pkgDir, "README.md")
	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}

func generateSidebar(result *analyzer.AnalyzerResult, cfg GeneratorConfig) error {
	var b strings.Builder
	b.WriteString("## Navigation\n\n")
	b.WriteString("- [🏠 Home](../README.md)\n")

	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			docPath := filepath.Join(pkg.Name, "README.md")
			b.WriteString(fmt.Sprintf("- [📦 %s](%s)\n", pkg.Name, docPath))
		}
	}

	return os.WriteFile(filepath.Join(cfg.OutputDir, "NAVIGATION.md"), []byte(b.String()), 0644)
}

func calculateDocCompletion(pkg analyzer.Package) int {
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

// processAIRequests assigns Documentation results to targets
func processAIRequests(requests []AIRequest, client *ai.Client) {
	ctx := context.Background()

	inputs := make([]string, len(requests))
	languages := make([]string, len(requests))
	for i, req := range requests {
		inputs[i] = req.Input
		languages[i] = req.Language
	}

	results, err := client.EnhanceDocumentationBatch(ctx, inputs, languages)
	if err != nil {
		fmt.Printf("⚠️ AI batch enhancement failed: %v\n", err)
		return
	}

	for i, res := range results {
		if requests[i].Target != nil {
			*requests[i].Target = res
		}
	}
}

func formatParams(params []analyzer.Parameter) string {
	var parts []string
	for _, p := range params {
		parts = append(parts, fmt.Sprintf("%s %s", p.Name, p.Type))
	}
	return strings.Join(parts, ", ")
}

func formatStruct(s analyzer.Struct) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("type %s struct {\n", s.Name))
	for _, f := range s.Fields {
		b.WriteString(fmt.Sprintf("\t%s %s %s\n", f.Name, f.Type, f.Tag))
	}
	b.WriteString("}")
	return b.String()
}

func formatFunction(f analyzer.Function) string {
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

// formatDocumentation formats a full Documentation struct to Markdown
func formatDocumentation(doc ai.Documentation) string {
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

func getLanguage(path string) string {
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
