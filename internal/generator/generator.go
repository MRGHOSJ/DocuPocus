package generator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	aiTypes "github.com/MRGHOSJ/docupocus/internal/ai/types"
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

type AICodeRequest struct {
	Input    string
	Language string // e.g. "Go", "Python"
	Target   *aiTypes.Documentation
}

type AIYAMLRequest struct {
	Input    string
	Language string // "YAML"
	Target   *aiTypes.YAMLDocumentation
}

func GeneratePackageDocs(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig) error {
	fmt.Println("📁 Creating output directory...")
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create doc directory: %w", err)
	}

	fmt.Println("📄 Generating project README...")
	if err := generateProjectReadme(result, cfg); err != nil {
		return fmt.Errorf("failed to generate project README: %w", err)
	}

	fmt.Println("📚 Generating sidebar...")
	if err := generateSidebar(result, cfg); err != nil {
		return fmt.Errorf("failed to generate sidebar: %w", err)
	}

	var codeRequests []AICodeRequest
	var yamlRequests []AIYAMLRequest

	fmt.Println("🔍 Preparing AI enhancement requests...")
	for fi := range result.Files {
		file := result.Files[fi]
		fmt.Printf("📦 File: %s\n", file.Path)

		for pi := range file.Packages {
			pkg := &file.Packages[pi]
			lang := getLanguage(file.Path)
			fmt.Printf("  📚 Package: %s (Lang: %s)\n", pkg.Name, lang)

			for si := range pkg.Structs {
				if cfg.AIClient == nil {
					continue
				}

				if lang == "YAML" {
					input := formatYAMLStruct(pkg.Structs[si])
					pkg.Structs[si].DocYAML = aiTypes.YAMLDocumentation{}
					yamlRequests = append(yamlRequests, AIYAMLRequest{
						Input:    input,
						Language: lang,
						Target:   &pkg.Structs[si].DocYAML,
					})
					fmt.Printf("    📄 YAML Struct: %s → YAML AI request added\n", pkg.Structs[si].Name)
				} else {
					input := formatStruct(pkg.Structs[si])
					pkg.Structs[si].Doc = aiTypes.Documentation{}
					codeRequests = append(codeRequests, AICodeRequest{
						Input:    input,
						Language: lang,
						Target:   &pkg.Structs[si].Doc,
					})
					fmt.Printf("    🧩 Struct: %s → Code AI request added\n", pkg.Structs[si].Name)
				}
			}

			if lang != "YAML" {
				for fi := range pkg.Funcs {
					if cfg.AIClient != nil {
						input := formatFunction(pkg.Funcs[fi])
						pkg.Funcs[fi].Doc = aiTypes.Documentation{}
						codeRequests = append(codeRequests, AICodeRequest{
							Input:    input,
							Language: lang,
							Target:   &pkg.Funcs[fi].Doc,
						})
						fmt.Printf("    🔧 Function: %s → Code AI request added\n", pkg.Funcs[fi].Name)
					}

				}
			}
		}
	}

	// Process requests separately
	if cfg.AIClient != nil {
		if len(codeRequests) > 0 || len(yamlRequests) > 0 {
			fmt.Printf("🚀 Sending %d code + %d YAML requests to AI\n", len(codeRequests), len(yamlRequests))
			processAIRequests(codeRequests, yamlRequests, cfg.AIClient)
			fmt.Println("✅ AI enhancement complete.")
		}
	}

	fmt.Println("📝 Generating documentation output...")

	// Generate final docs
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			if getLanguage(file.Path) == "YAML" {
				fmt.Printf("📄 Generating YAML documentation for: %s\n", pkg.Name)
				if err := generateYAMLDoc(pkg, file.Path, cfg); err != nil {
					return fmt.Errorf("failed to generate YAML docs for package %s: %w", pkg.Name, err)
				}
			} else {
				fmt.Printf("📄 Generating code documentation for: %s\n", pkg.Name)
				if err := generatePackageDoc(pkg, file.Path, cfg); err != nil {
					return fmt.Errorf("failed to generate docs for package %s: %w", pkg.Name, err)
				}
			}
		}
	}

	fmt.Println("✅ All documentation successfully generated.")
	return nil
}

func generateYAMLDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig) error {
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
	b.WriteString(fmt.Sprintf("> 📍 Path: `%s`\n\n", getDisplayPath(filePath)))

	for _, s := range pkg.Structs {
		doc := s.DocYAML

		// Expandable Resource Summary
		if doc.Summary != "" {
			b.WriteString("<details>\n")
			b.WriteString("<summary>🚀 Resource Summary</summary>\n\n")
			b.WriteString(fmt.Sprintf("- **Kind:** `%s`\n", s.Name))
			b.WriteString(fmt.Sprintf("- **Description:** %s\n\n", doc.Summary))
			b.WriteString("</details>\n\n")
		}

		// Expandable Configuration Example
		if len(s.Fields) > 0 {
			b.WriteString("<details>\n")
			b.WriteString(fmt.Sprintf("<summary>⚙️ Configuration Example for `%s`</summary>\n\n", s.Name))
			b.WriteString("```yaml\n")
			b.WriteString(generateYAMLExample(s.Fields, 0))
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
	indent := strings.Repeat("  ", indentLevel)

	for _, f := range fields {
		fieldName := f.Name
		fieldVal := f.Value

		switch f.Type {
		case "array":
			b.WriteString(fmt.Sprintf("%s%s:\n", indent, fieldName))
			for _, item := range f.Fields {
				innerIndent := strings.Repeat("  ", indentLevel+1)
				b.WriteString(innerIndent + "- ")
				if len(item.Fields) > 0 {
					first := item.Fields[0]
					b.WriteString(fmt.Sprintf("%s: %s\n", first.Name, first.Value))
					for _, sub := range item.Fields[1:] {
						b.WriteString(fmt.Sprintf("%s  %s: %s\n", innerIndent, sub.Name, sub.Value))
					}
				} else {
					b.WriteString(item.Value + "\n")
				}
			}

		case "map":
			b.WriteString(fmt.Sprintf("%s%s:\n", indent, fieldName))
			b.WriteString(generateYAMLExample(f.Fields, indentLevel+1))

		default:
			if fieldVal != "" {
				b.WriteString(fmt.Sprintf("%s%s: %s\n", indent, fieldName, fieldVal))
			} else {
				b.WriteString(fmt.Sprintf("%s%s: # TODO: Add value\n", indent, fieldName))
			}
		}
	}

	return b.String()
}

// Helper to clean up empty markdown cells
func safeMarkdown(s string) string {
	if s == "" {
		return "N/A"
	}
	return s
}

func isServiceFile(pkg analyzer.Package) bool {
	for _, s := range pkg.Structs {
		if strings.ToLower(s.Name) == "service" {
			return true
		}
	}
	return false
}

// Helper to check if this is a Kubernetes YAML file
func isKubernetesFile(pkg analyzer.Package) bool {
	for _, s := range pkg.Structs {
		if strings.Contains(s.Doc.Summary, "Kubernetes") {
			return true
		}
	}
	return false
}

// New function to format YAML structs for AI processing
func formatYAMLStruct(s analyzer.Struct) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("YAML Configuration Structure: %s\n", s.Name))
	if s.Doc.Summary != "" {
		b.WriteString(fmt.Sprintf("Description: %s\n", s.Doc.Summary))
	}
	b.WriteString("Fields:\n")

	for _, f := range s.Fields {
		b.WriteString(fmt.Sprintf("- %s (%s)", f.Name, f.Type))
		if f.Value != "" {
			b.WriteString(fmt.Sprintf(" = %s", f.Value))
		}
		b.WriteString("\n")

		if len(f.Fields) > 0 {
			b.WriteString(formatYAMLFields(f.Fields, 1))
		}
	}

	return b.String()
}

// Helper for recursive field formatting
func formatYAMLFields(fields []analyzer.Field, depth int) string {
	var b strings.Builder
	indent := strings.Repeat("  ", depth)

	for _, f := range fields {
		b.WriteString(indent + fmt.Sprintf("- %s (%s)", f.Name, f.Type))
		if f.Value != "" {
			b.WriteString(fmt.Sprintf(" = %s", f.Value))
		}
		b.WriteString("\n")

		if len(f.Fields) > 0 {
			b.WriteString(formatYAMLFields(f.Fields, depth+1))
		}
	}

	return b.String()
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
	b.WriteString(fmt.Sprintf("> 📍 `%s`\n\n", getDisplayPath(filePath)))

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

	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}

func getDisplayPath(fullPath string) string {
	parts := strings.Split(filepath.Clean(fullPath), string(filepath.Separator))
	if len(parts) < 2 {
		return fullPath
	}
	return filepath.Join(parts[len(parts)-2:]...)
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
func processAIRequests(
	codeRequests []AICodeRequest,
	yamlRequests []AIYAMLRequest,
	client *ai.Client,
) {
	ctx := context.Background()

	// Process code requests
	if len(codeRequests) > 0 {
		inputs := make([]string, len(codeRequests))
		languages := make([]string, len(codeRequests))
		for i, req := range codeRequests {
			inputs[i] = req.Input
			languages[i] = req.Language
		}

		results, err := client.EnhanceDocumentationBatch(ctx, inputs, languages)
		if err != nil {
			fmt.Printf("⚠️ Code enhancement failed: %v\n", err)
		} else {
			for i, res := range results {
				if codeRequests[i].Target != nil {
					*codeRequests[i].Target = res
				}
			}
		}
	}

	// Process YAML requests
	if len(yamlRequests) > 0 {
		inputs := make([]string, len(yamlRequests))
		languages := make([]string, len(yamlRequests))
		for i, req := range yamlRequests {
			inputs[i] = req.Input
			languages[i] = req.Language
		}

		results, err := client.EnhanceYAMLDocumentationBatch(ctx, inputs, languages)
		if err != nil {
			fmt.Printf("⚠️ YAML enhancement failed: %v\n", err)
		} else {
			for i, res := range results {
				if yamlRequests[i].Target != nil {
					*yamlRequests[i].Target = res
				}
			}
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
