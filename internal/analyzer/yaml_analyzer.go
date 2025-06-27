package analyzer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ai "github.com/MRGHOSJ/docupocus/internal/ai/types"
	"github.com/MRGHOSJ/docupocus/internal/utils"
	"gopkg.in/yaml.v3"
)

type YAMLAnalyzer struct{}

func (y *YAMLAnalyzer) Supports(projectDir string) bool {
	return utils.HasFilesWithExtension(projectDir, ".yaml") ||
		utils.HasFilesWithExtension(projectDir, ".yml")
}

func (y *YAMLAnalyzer) Analyze(projectDir string) (*AnalyzerResult, error) {
	var result AnalyzerResult

	err := filepath.WalkDir(projectDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !(strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml")) {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Parse with full document structure
		var nodes []yaml.Node
		decoder := yaml.NewDecoder(strings.NewReader(string(content)))
		for {
			var node yaml.Node
			if err := decoder.Decode(&node); err != nil {
				break
			}
			nodes = append(nodes, node)
		}

		pkg := Package{
			Name: filepath.Base(filepath.Dir(path)),
			Path: path,
		}

		// Handle multi-document YAML files
		for _, node := range nodes {
			structs := y.analyzeYAMLNode(&node, filepath.Base(path))
			pkg.Structs = append(pkg.Structs, structs...)
		}

		result.Files = append(result.Files, &AnalyzedFile{
			Path:     path,
			Packages: []Package{pkg},
		})

		return nil
	})

	return &result, err
}

func (y *YAMLAnalyzer) analyzeYAMLNode(node *yaml.Node, baseName string) []Struct {
	var structs []Struct

	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		node = node.Content[0]
	}

	// Handle different YAML document structures
	switch {
	case y.isKubernetesResource(node):
		resource := y.analyzeKubernetesResource(node, baseName)
		resource.Doc.Summary = y.extractYAMLDescription(node)
		structs = append(structs, resource)
	case y.isHelmChart(node):
		chart := y.analyzeHelmChart(node, baseName)
		chart.Doc.Summary = y.extractYAMLDescription(node)
		structs = append(structs, chart)
	case y.isAnsiblePlaybook(node):
		chart := y.analyzeAnsiblePlaybook(node, baseName)
		chart.Doc.Summary = y.extractYAMLDescription(node)
		structs = append(structs, chart)
	case y.isDockerCompose(node):
		chart := y.analyzeDockerCompose(node, baseName)
		chart.Doc.Summary = y.extractYAMLDescription(node)
		structs = append(structs, chart)
	default:
		// Generic YAML analysis
		structs = append(structs, Struct{
			Name:    baseName,
			DocYAML: ai.YAMLDocumentation{Summary: y.extractYAMLDescription(node)},
			Fields:  y.extractYAMLFields(node, 0),
		})
	}

	return structs
}

// -- Specialized YAML Type Detectors --

func (y *YAMLAnalyzer) isKubernetesResource(node *yaml.Node) bool {
	return node.Kind == yaml.MappingNode &&
		y.hasKey(node, "apiVersion") &&
		y.hasKey(node, "kind")
}

func (y *YAMLAnalyzer) isHelmChart(node *yaml.Node) bool {
	return node.Kind == yaml.MappingNode &&
		(y.hasKey(node, "chart") ||
			y.hasKey(node, "dependencies"))
}

func (y *YAMLAnalyzer) isAnsiblePlaybook(node *yaml.Node) bool {
	return node.Kind == yaml.SequenceNode &&
		len(node.Content) > 0 &&
		node.Content[0].Kind == yaml.MappingNode &&
		y.hasKey(node.Content[0], "hosts")
}

func (y *YAMLAnalyzer) isDockerCompose(node *yaml.Node) bool {
	return node.Kind == yaml.MappingNode &&
		(y.hasKey(node, "services") ||
			y.hasKey(node, "version"))
}

// -- Specialized Analyzers --

func (y *YAMLAnalyzer) analyzeKubernetesResource(node *yaml.Node, baseName string) Struct {
	kind := y.getValue(node, "kind")
	if kind == "" {
		kind = baseName
	}

	// Get metadata for better description
	metadata := y.getValueNode(node, "metadata")
	name := ""
	if metadata != nil {
		name = y.getValue(metadata, "name")
	}

	desc := fmt.Sprintf("Kubernetes %s", kind)
	if name != "" {
		desc += fmt.Sprintf(": %s", name)
	}
	if v := y.getValue(node, "apiVersion"); v != "" {
		desc += fmt.Sprintf(" (apiVersion: %s)", v)
	}

	return Struct{
		Name:    kind,
		DocYAML: ai.YAMLDocumentation{Summary: desc},
		Fields:  y.extractKubernetesSpec(node),
	}
}

func (y *YAMLAnalyzer) extractKubernetesSpec(node *yaml.Node) []Field {
	var fields []Field

	// Extract fields at the root level
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 >= len(node.Content) {
				continue
			}
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]

			// // Skip known non-config sections
			// if keyNode.Value == "apiVersion" || keyNode.Value == "kind" || keyNode.Value == "metadata" {
			// 	continue
			// }

			field := Field{
				Name:    keyNode.Value,
				Type:    y.nodeKindToString(valueNode.Kind),
				DocYAML: ai.YAMLDocumentation{Summary: strings.TrimSpace(keyNode.HeadComment + "\n" + keyNode.LineComment)},
			}

			if valueNode.Kind == yaml.ScalarNode {
				field.Value = valueNode.Value
			}
			if valueNode.Kind == yaml.MappingNode || valueNode.Kind == yaml.SequenceNode {
				field.Fields = y.extractYAMLFields(valueNode, 1)
			}

			fields = append(fields, field)
		}
	}

	return fields
}

func (y *YAMLAnalyzer) findSpecNode(node *yaml.Node) *yaml.Node {
	if node.Kind != yaml.MappingNode {
		return nil
	}

	// Look for 'spec' key
	for i := 0; i < len(node.Content); i += 2 {
		if i+1 >= len(node.Content) {
			continue
		}
		keyNode := node.Content[i]
		if keyNode.Value == "spec" {
			return node.Content[i+1]
		}
	}

	return nil
}

func (y *YAMLAnalyzer) analyzeHelmChart(node *yaml.Node, baseName string) Struct {
	return Struct{
		Name: "HelmChart",
		DocYAML: ai.YAMLDocumentation{
			Summary: "Helm Chart configuration \n" + y.extractYAMLDescription(node),
		},
		Fields: y.extractYAMLFields(node, 0),
	}
}

func (y *YAMLAnalyzer) analyzeAnsiblePlaybook(node *yaml.Node, baseName string) Struct {
	playbook := Struct{
		Name: "AnsiblePlaybook",
		DocYAML: ai.YAMLDocumentation{
			Summary: "Ansible Playbook configuration",
		},
	}

	if node.Kind == yaml.SequenceNode && len(node.Content) > 0 {
		// Analyze first play (common case)
		firstPlay := node.Content[0]
		if firstPlay.Kind == yaml.MappingNode {
			playbook.Fields = y.extractAnsiblePlayFields(firstPlay)
		}
	}

	return playbook
}

func (y *YAMLAnalyzer) extractAnsiblePlayFields(playNode *yaml.Node) []Field {
	var fields []Field

	for i := 0; i < len(playNode.Content); i += 2 {
		if i+1 >= len(playNode.Content) {
			continue
		}
		keyNode := playNode.Content[i]
		valueNode := playNode.Content[i+1]

		field := Field{
			Name:    keyNode.Value,
			Type:    y.nodeKindToString(valueNode.Kind),
			DocYAML: ai.YAMLDocumentation{Summary: strings.TrimSpace(keyNode.HeadComment + "\n" + keyNode.LineComment)},
		}

		// Special handling for tasks
		if keyNode.Value == "tasks" && valueNode.Kind == yaml.SequenceNode {
			field.Fields = y.extractAnsibleTasks(valueNode)
		} else if valueNode.Kind == yaml.MappingNode || valueNode.Kind == yaml.SequenceNode {
			field.Fields = y.extractYAMLFields(valueNode, 1)
		}

		fields = append(fields, field)
	}

	return fields
}

func (y *YAMLAnalyzer) extractAnsibleTasks(tasksNode *yaml.Node) []Field {
	var taskFields []Field

	for i, taskNode := range tasksNode.Content {
		if taskNode.Kind == yaml.MappingNode {
			taskField := Field{
				Name:   "task_" + string(rune('a'+i)),
				Type:   "task",
				Fields: []Field{},
			}

			// Extract task details
			for j := 0; j < len(taskNode.Content); j += 2 {
				if j+1 >= len(taskNode.Content) {
					continue
				}
				keyNode := taskNode.Content[j]
				valueNode := taskNode.Content[j+1]

				subField := Field{
					Name:    keyNode.Value,
					Type:    y.nodeKindToString(valueNode.Kind),
					DocYAML: ai.YAMLDocumentation{Summary: strings.TrimSpace(keyNode.HeadComment + "\n" + keyNode.LineComment)},
				}

				if valueNode.Kind == yaml.MappingNode || valueNode.Kind == yaml.SequenceNode {
					subField.Fields = y.extractYAMLFields(valueNode, 2)
				}

				taskField.Fields = append(taskField.Fields, subField)
			}

			taskFields = append(taskFields, taskField)
		}
	}

	return taskFields
}

func (y *YAMLAnalyzer) analyzeDockerCompose(node *yaml.Node, baseName string) Struct {
	compose := Struct{
		Name: "DockerCompose",
		DocYAML: ai.YAMLDocumentation{
			Summary: "Docker Compose configuration",
		},
	}

	// Extract version if present
	version := y.getValue(node, "version")
	if version != "" {
		compose.Doc.Summary += "\nVersion: " + version
	}

	// Extract services
	servicesNode := y.getValueNode(node, "services")
	if servicesNode != nil && servicesNode.Kind == yaml.MappingNode {
		compose.Fields = y.extractDockerServices(servicesNode)
	}

	// Extract networks and volumes if present
	if networksNode := y.getValueNode(node, "networks"); networksNode != nil {
		compose.Fields = append(compose.Fields, Field{
			Name:   "networks",
			Type:   y.nodeKindToString(networksNode.Kind),
			Fields: y.extractYAMLFields(networksNode, 1),
		})
	}

	if volumesNode := y.getValueNode(node, "volumes"); volumesNode != nil {
		compose.Fields = append(compose.Fields, Field{
			Name:   "volumes",
			Type:   y.nodeKindToString(volumesNode.Kind),
			Fields: y.extractYAMLFields(volumesNode, 1),
		})
	}

	return compose
}

func (y *YAMLAnalyzer) extractDockerServices(servicesNode *yaml.Node) []Field {
	var serviceFields []Field

	for i := 0; i < len(servicesNode.Content); i += 2 {
		if i+1 >= len(servicesNode.Content) {
			continue
		}
		serviceNameNode := servicesNode.Content[i]
		serviceConfigNode := servicesNode.Content[i+1]

		serviceField := Field{
			Name:    serviceNameNode.Value,
			Type:    "service",
			DocYAML: ai.YAMLDocumentation{Summary: strings.TrimSpace(serviceNameNode.HeadComment + "\n" + serviceNameNode.LineComment)},
		}

		if serviceConfigNode.Kind == yaml.MappingNode {
			serviceField.Fields = y.extractYAMLFields(serviceConfigNode, 1)
		}

		serviceFields = append(serviceFields, serviceField)
	}

	return serviceFields
}

// Helper to get the value node for a key
func (y *YAMLAnalyzer) getValueNode(node *yaml.Node, key string) *yaml.Node {
	if node.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i < len(node.Content); i += 2 {
		if i+1 < len(node.Content) && node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

// -- Core Analysis Helpers --

func (y *YAMLAnalyzer) extractYAMLFields(node *yaml.Node, depth int) []Field {
	var fields []Field

	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 >= len(node.Content) {
				continue
			}
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]

			field := Field{
				Name:    keyNode.Value,
				Type:    y.nodeKindToString(valueNode.Kind),
				DocYAML: ai.YAMLDocumentation{Summary: strings.TrimSpace(keyNode.HeadComment + "\n" + keyNode.LineComment)},
			}

			// Capture scalar values
			if valueNode.Kind == yaml.ScalarNode {
				field.Value = valueNode.Value
			}

			// Handle nested structures
			if valueNode.Kind == yaml.MappingNode || valueNode.Kind == yaml.SequenceNode {
				field.Fields = y.extractYAMLFields(valueNode, depth+1)
			}

			fields = append(fields, field)
		}
	} else if node.Kind == yaml.SequenceNode {
		for i, itemNode := range node.Content {
			field := Field{
				Name:   fmt.Sprintf("[%d]", i),
				Type:   y.nodeKindToString(itemNode.Kind),
				Fields: y.extractYAMLFields(itemNode, depth+1),
			}
			fields = append(fields, field)
		}
	}

	return fields
}

func (y *YAMLAnalyzer) extractYAMLDescription(node *yaml.Node) string {
	var desc strings.Builder

	if node.HeadComment != "" {
		desc.WriteString(strings.TrimSpace(node.HeadComment))
		desc.WriteString("\n\n")
	}

	switch node.Kind {
	case yaml.MappingNode:
		desc.WriteString("YAML mapping with keys: ")
		keys := y.getMappingKeys(node)
		desc.WriteString(strings.Join(keys, ", "))
	case yaml.SequenceNode:
		desc.WriteString("YAML sequence with ")
		desc.WriteString(string(len(node.Content)))
		desc.WriteString(" items")
	case yaml.ScalarNode:
		desc.WriteString("YAML scalar value")
	}

	return desc.String()
}

// -- Utility Methods --

func (y *YAMLAnalyzer) hasKey(node *yaml.Node, key string) bool {
	if node.Kind != yaml.MappingNode {
		return false
	}
	for i := 0; i < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return true
		}
	}
	return false
}

func (y *YAMLAnalyzer) getValue(node *yaml.Node, key string) string {
	if node.Kind != yaml.MappingNode {
		return ""
	}
	for i := 0; i < len(node.Content); i += 2 {
		if node.Content[i].Value == key && i+1 < len(node.Content) {
			return node.Content[i+1].Value
		}
	}
	return ""
}

func (y *YAMLAnalyzer) getMappingKeys(node *yaml.Node) []string {
	var keys []string
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			keys = append(keys, node.Content[i].Value)
		}
	}
	return keys
}

func (y *YAMLAnalyzer) nodeKindToString(kind yaml.Kind) string {
	switch kind {
	case yaml.DocumentNode:
		return "document"
	case yaml.SequenceNode:
		return "array"
	case yaml.MappingNode:
		return "map"
	case yaml.ScalarNode:
		return "scalar"
	case yaml.AliasNode:
		return "alias"
	default:
		return "unknown"
	}
}

func (y *YAMLAnalyzer) sequenceItemName(index int) string {
	return "item_" + string(rune('a'+index))
}
