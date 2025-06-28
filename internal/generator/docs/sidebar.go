package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	docTypes "github.com/MRGHOSJ/docupocus/internal/generator/types"
)

func GenerateSidebar(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig) error {
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
