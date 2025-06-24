package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"
)

type GeneratorConfig struct {
	AIClient *ai.Client
}

// AIRequest now targets a full Documentation struct
type AIRequest struct {
	Input    string
	Language string
	Target   *ai.Documentation
}

func GenerateMarkdown(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig) (string, error) {
	var b strings.Builder

	// Step 1: Generate project title and description using AI
	if cfg.AIClient != nil {
		projectSummary := struct {
			Title       string
			Description string
		}{
			Title:       "Project Title",
			Description: "Project Description",
		}
		b.WriteString(fmt.Sprintf("# 📘 %s\n\n%s\n\n", projectSummary.Title, projectSummary.Description))
	}

	// Step 2: Packages Overview
	b.WriteString("## 📦 Packages Overview\n\n")
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			b.WriteString(fmt.Sprintf("- **%s** in _%s_\n", pkg.Name, file.Path))
		}
	}
	b.WriteString("\n")

	// Step 3: Collect AI requests for code-level docs
	var aiRequests []AIRequest
	for fi := range result.Files {
		for pi := range result.Files[fi].Packages {
			pkg := &result.Files[fi].Packages[pi]
			lang := getLanguage(result.Files[fi].Path)

			for si := range pkg.Structs {
				if cfg.AIClient != nil {
					input := formatStruct(pkg.Structs[si])
					pkg.Structs[si].Doc = ai.Documentation{} // Ensure it's initialized
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

	// Step 4: Enhance documentation using AI
	if cfg.AIClient != nil && len(aiRequests) > 0 {
		fmt.Printf("🚀 Sending %d documentation enhancements to AI\n", len(aiRequests))
		processAIRequests(aiRequests, cfg.AIClient)
	}

	// Step 5: Core Package Documentation
	b.WriteString("## 🛠️ Core Packages\n\n")
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			b.WriteString(fmt.Sprintf("### 📦 Package: `%s`\n\n", pkg.Name))

			if len(pkg.Structs) > 0 {
				b.WriteString("#### 🧱 Structs:\n\n")
				for _, s := range pkg.Structs {
					b.WriteString(fmt.Sprintf("- **%s**\n", s.Name))
					b.WriteString(formatDocumentation(s.Doc))
				}
				b.WriteString("\n")
			}

			if len(pkg.Funcs) > 0 {
				b.WriteString("#### 🔧 Functions:\n\n")
				for _, f := range pkg.Funcs {
					b.WriteString(fmt.Sprintf("- **%s(%s)**\n", f.Name, formatParams(f.Parameters)))
					b.WriteString(formatDocumentation(f.Doc))
				}
				b.WriteString("\n")
			}
		}
	}

	// Step 6: Diagrams section (placeholder)
	b.WriteString("## 📊 Diagrams\n\n_Generated architecture or flow diagrams will appear here._\n\n")

	// Step 7: Best Practices
	b.WriteString("## 💡 Best Practices\n\n")
	b.WriteString("- Keep functions short and descriptive.\n")
	b.WriteString("- Write meaningful comments for exported symbols.\n")
	b.WriteString("- Avoid cyclic dependencies across packages.\n\n")

	// Step 8: Quick Start Guide
	b.WriteString("## 🚀 Quick Start\n\n")
	b.WriteString("```bash\n# Clone the repository\ngit clone https://github.com/your/project.git\n\n# Build\ngo build ./...\n\n# Run\ngo run main.go\n```\n\n")

	return b.String(), nil
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
