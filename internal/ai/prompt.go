package ai

import (
	"fmt"
	"strings"
)

func (c *Client) buildPrompt(input, language string) string {
	if c.config.Prompt == "" {
		return fmt.Sprintf("Generate a concise explanation (10-15 words) for this %s code:\n```%s\n%s\n```",
			language, language, input)
	}
	return fmt.Sprintf(c.config.Prompt, language, language, input)
}

func (c *Client) buildBatchPromptYamlDocumentation(prompts []string) string {
	var sb strings.Builder

	sb.WriteString("You are an **advanced YAML configuration documentation assistant**.\n\n")
	sb.WriteString("For each YAML snippet, return a **JSON object** with the following fields:\n")
	sb.WriteString("- summary: brief purpose (max 20 words)\n")
	sb.WriteString("- fields: list of key fields with name, type (scalar, map, array), and short description\n")
	sb.WriteString("- examples: valid example values for key fields (if applicable)\n")
	sb.WriteString("- defaults: known default values\n")
	sb.WriteString("- usage: common usage scenario\n")
	sb.WriteString("- best_practices: warnings, constraints, or best practices\n\n")
	sb.WriteString("Return a **JSON array**, one object per snippet, like:\n\n")

	sb.WriteString(`[
  {
    "summary": "Describes access modes and storage class for a volume",
    "fields": [
      { "name": "accessModes", "type": "array", "description": "Mount options: ReadWriteOnce, etc." },
      { "name": "storageClassName", "type": "scalar", "description": "Storage class name" }
    ],
    "examples": {
      "accessModes": ["ReadWriteOnce"],
      "storageClassName": "standard"
    },
    "defaults": {
      "storageClassName": "default"
    },
    "usage": "Used in PersistentVolumeClaim for storage provisioning",
    "best_practices": [
      "Explicitly define accessModes",
      "Use standard storage class names"
    ]
  }
]`)

	sb.WriteString("\n\nYAML snippets:\n[\n")
	for i, prompt := range prompts {
		cleanPrompt := strings.TrimSpace(prompt)
		escapedPrompt := strings.ReplaceAll(cleanPrompt, `"`, `\"`)
		sb.WriteString(fmt.Sprintf(`  "Snippet %d: %s"`, i+1, escapedPrompt))
		if i < len(prompts)-1 {
			sb.WriteString(",\n")
		}
	}
	sb.WriteString("\n]\n")

	return sb.String()
}

func (c *Client) buildBatchPromptCodeAssistant(prompts []string) string {
	var sb strings.Builder

	// Enhanced prompt with multiple documentation requirements
	sb.WriteString("You are an **advanced code documentation assistant**. For each code snippet provided, generate:\n\n")
	sb.WriteString("- A concise functional summary (15-20 words)\n")
	sb.WriteString("- Key parameters/inputs with types\n")
	sb.WriteString("- Return value/output description\n")
	sb.WriteString("- Time complexity analysis (Big-O notation)\n")
	sb.WriteString("- Space complexity analysis\n")
	sb.WriteString("- One common usage example\n")
	sb.WriteString("- Potential edge cases to consider\n\n")

	sb.WriteString(`**Return a JSON array** where each element is **an object containing these documentation aspects**.`)
	sb.WriteString("\n\nCode snippets:\n[\n")

	for i, prompt := range prompts {
		// Properly escape JSON and trim whitespace
		cleanPrompt := strings.TrimSpace(prompt)
		escapedPrompt := strings.ReplaceAll(cleanPrompt, `"`, `\"`)
		sb.WriteString(fmt.Sprintf(`  "Snippet %d: %s"`, i+1, escapedPrompt))
		if i < len(prompts)-1 {
			sb.WriteString(",\n")
		}
	}

	sb.WriteString("\n]\n\n")
	sb.WriteString(`**Return format example**:
[
  {
    "summary": "Function that adds two integers",
    "parameters": [
      {"name": "a", "type": "int", "description": "First operand"},
      {"name": "b", "type": "int", "description": "Second operand"}
    ],
    "returns": "Sum of the two integers as int",
    "time_complexity": "O(1)",
    "space_complexity": "O(1)",
    "usage_example": "sum := add(3, 5) // returns 8",
    "edge_cases": [
      "Integer overflow with large numbers",
      "Undefined behavior with extremely large values"
    ]
  }
]`)

	return sb.String()
}
