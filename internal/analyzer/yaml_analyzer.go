package analyzer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/utils"
	"gopkg.in/yaml.v3"
)

type YAMLAnalyzer struct{}

func (y *YAMLAnalyzer) Supports(projectDir string) bool {
	return utils.HasFilesWithExtension(projectDir, ".yaml") || utils.HasFilesWithExtension(projectDir, ".yml")
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

		// Parse YAML to extract key structure
		var node yaml.Node
		if err := yaml.Unmarshal(content, &node); err != nil {
			return err
		}

		pkg := Package{
			Name: filepath.Base(filepath.Dir(path)),
			Path: path,
			Structs: []Struct{
				{
					Name: filepath.Base(path),
					Doc:  ai.Documentation{Summary: extractYAMLDescription(&node)},
				},
			},
		}

		result.Files = append(result.Files, &AnalyzedFile{
			Path:     path,
			Packages: []Package{pkg},
		})

		return nil
	})

	return &result, err
}

func extractYAMLDescription(node *yaml.Node) string {
	// Simple extraction of top-level keys as description
	var keys []string
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		node = node.Content[0]
	}
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i].Value
			keys = append(keys, key)
		}
	}
	return "YAML configuration with keys: " + strings.Join(keys, ", ")
}
