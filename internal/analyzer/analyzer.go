package analyzer

import (
	"fmt"

	"github.com/MRGHOSJ/docupocus/internal/ai"
)

type Analyzer interface {
	Analyze(projectDir string) (*AnalyzerResult, error)
	Supports(projectDir string) bool
}

// AnalyzerResult holds the results of the analysis
type AnalyzerResult struct {
	Files []*AnalyzedFile
}

type AnalyzedFile struct {
	Path     string
	Packages []Package // Replace this with actual extracted structs/functions
}

type Package struct {
	Name    string
	Path    string
	Structs []Struct
	Funcs   []Function
	Files   []string
}

type Struct struct {
	Name    string
	Fields  []Field
	Methods []Function
	Doc     ai.Documentation
}

type Field struct {
	Name string
	Type string
	Tag  string
	Doc  ai.Documentation
}

type Function struct {
	Name       string
	Receiver   string
	Parameters []Parameter
	Results    []Parameter
	Doc        ai.Documentation
	Calls      []string
}

type Parameter struct {
	Name string
	Type string
}

var analyzers []Analyzer

func init() {
	analyzers = []Analyzer{
		&GoAnalyzer{},
		&PythonAnalyzer{},
		&JSAnalyzer{},
		&YAMLAnalyzer{},
		// More...
	}
}

func AnalyzeProject(projectDir string) (*AnalyzerResult, error) {
	for _, a := range analyzers {
		if a.Supports(projectDir) {
			return a.Analyze(projectDir)
		}
	}
	return nil, fmt.Errorf("no analyzer found for project in: %s", projectDir)
}
