package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	cfg "github.com/MRGHOSJ/docupocus/internal/generator/types"
)

func processAIRequests(
	codeRequests []cfg.AICodeRequest,
	yamlRequests []cfg.AIYAMLRequest,
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
