package analyzer

import (
	"fmt"
	"path/filepath"
	"strings"

	ai "github.com/MRGHOSJ/docupocus/internal/ai/types"
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
	Imports []string
	Structs []Struct
	Funcs   []Function
	Files   []string
}

type Struct struct {
	Name    string
	Fields  []Field
	Methods []Function
	Doc     ai.Documentation
	DocYAML ai.YAMLDocumentation
}

type Field struct {
	Name    string
	Type    string
	Tag     string
	Doc     ai.Documentation
	DocYAML ai.YAMLDocumentation
	Value   string
	Fields  []Field
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
		analyzerName := strings.TrimPrefix(fmt.Sprintf("%T", a), "*analyzer.")
		fmt.Printf("🔍 Checking with %s analyzer... ", analyzerName)

		if a.Supports(projectDir) {
			fmt.Printf("MATCH\n")
			result, err := a.Analyze(projectDir)
			if err != nil {
				fmt.Printf("❌ Analysis failed: %v\n", err)
				return nil, err
			}

			fmt.Printf("\n📦 Analysis Results (%d files):\n", len(result.Files))
			for _, file := range result.Files {
				yamlPrintFileAnalysis(file)
			}
			return result, nil
		} else {
			fmt.Printf("skip\n")
		}
	}
	return nil, fmt.Errorf("no analyzer found for project in: %s", projectDir)
}
func yamlPrintFileAnalysis(file *AnalyzedFile) {
	fmt.Printf("\n📄 %s\n", filepath.Base(file.Path))

	for _, pkg := range file.Packages {
		for _, s := range pkg.Structs {
			// Print resource type and basic info
			fmt.Printf("  🏷️ %s\n", s.Name)
			if s.Doc.Summary != "" {
				fmt.Printf("  📝 %s\n", s.Doc.Summary)
			}

			// Print important fields with values
			if len(s.Fields) > 0 {
				fmt.Println("  ⚙️ Configuration:")
				for _, f := range s.Fields {
					yamlPrintField(f, 4)
				}
			}
			fmt.Println("  ──────────────────────")
		}
	}
}

func yamlPrintField(f Field, indent int) {
	indentStr := strings.Repeat(" ", indent)

	// Print field name and type
	fmt.Printf("%s- %s (%s)", indentStr, f.Name, f.Type)

	// Print value if available
	if f.Value != "" {
		fmt.Printf(" = %s", f.Value)
	}
	fmt.Println()

	// Print documentation if available
	if f.Doc.Summary != "" {
		fmt.Printf("%s  # %s\n", indentStr, f.Doc.Summary)
	}

	// Recursively print nested fields
	for _, subField := range f.Fields {
		yamlPrintField(subField, indent+2)
	}
}
