package generator

import (
	"github.com/MRGHOSJ/docupocus/internal/ai"
	aiTypes "github.com/MRGHOSJ/docupocus/internal/ai/types"
)

type GeneratorConfig struct {
	AIClient  *ai.Client
	OutputDir string
	Project   ProjectMeta
}

type ProjectMeta struct {
	Name          string
	Description   string
	RepoURL       string
	Features      []Feature
	TechStack     []string
	QuickStart    []QuickStartBlock
	BestPractices BestPractices
}

type Feature struct {
	Title       string
	Description string
}

type QuickStartBlock struct {
	Title   string // e.g. "▶️ Local"
	Shell   string // e.g. "bash"
	Command string // full command block
}

type BestPractices struct {
	Do   []string
	Dont []string
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
