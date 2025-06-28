package generator

import (
	"fmt"
	"os"

	aiTypes "github.com/MRGHOSJ/docupocus/internal/ai/types"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"

	docGenerator "github.com/MRGHOSJ/docupocus/internal/generator/docs"
	docTypes "github.com/MRGHOSJ/docupocus/internal/generator/types"
	docUtils "github.com/MRGHOSJ/docupocus/internal/generator/utils"
)

func GeneratePackageDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig) error {
	if err := prepareOutputStructure(result, cfg); err != nil {
		return err
	}

	codeRequests, yamlRequests := prepareAIRequests(result, cfg)

	if err := enhanceWithAI(codeRequests, yamlRequests, cfg); err != nil {
		return err
	}

	return generateFinalDocs(result, cfg)
}

func prepareOutputStructure(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig) error {
	fmt.Println("📁 Creating output directory...")
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create doc directory: %w", err)
	}

	fmt.Println("📄 Generating project README...")
	if err := docGenerator.GenerateProjectReadme(result, cfg); err != nil {
		return fmt.Errorf("failed to generate project README: %w", err)
	}

	fmt.Println("📚 Generating sidebar...")
	if err := docGenerator.GenerateSidebar(result, cfg); err != nil {
		return fmt.Errorf("failed to generate sidebar: %w", err)
	}

	return nil
}

func prepareAIRequests(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig) ([]docTypes.AICodeRequest, []docTypes.AIYAMLRequest) {
	var codeRequests []docTypes.AICodeRequest
	var yamlRequests []docTypes.AIYAMLRequest

	fmt.Println("🔍 Preparing AI enhancement requests...")
	for fi := range result.Files {
		file := result.Files[fi]
		fmt.Printf("📦 File: %s\n", file.Path)

		for pi := range file.Packages {
			pkg := &file.Packages[pi]
			lang := docUtils.GetLanguage(file.Path)
			fmt.Printf("  📚 Package: %s (Lang: %s)\n", pkg.Name, lang)

			for si := range pkg.Structs {
				if cfg.AIClient == nil {
					continue
				}
				if lang == "YAML" {
					input := formatYAMLStruct(pkg.Structs[si])
					pkg.Structs[si].DocYAML = aiTypes.YAMLDocumentation{}
					yamlRequests = append(yamlRequests, docTypes.AIYAMLRequest{
						Input:    input,
						Language: lang,
						Target:   &pkg.Structs[si].DocYAML,
					})
					fmt.Printf("    📄 YAML Struct: %s → YAML AI request added\n", pkg.Structs[si].Name)
				} else {
					input := docGenerator.FormatStruct(pkg.Structs[si])
					pkg.Structs[si].Doc = aiTypes.Documentation{}
					codeRequests = append(codeRequests, docTypes.AICodeRequest{
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
						input := docGenerator.FormatFunction(pkg.Funcs[fi])
						pkg.Funcs[fi].Doc = aiTypes.Documentation{}
						codeRequests = append(codeRequests, docTypes.AICodeRequest{
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
	return codeRequests, yamlRequests
}

func enhanceWithAI(codeRequests []docTypes.AICodeRequest, yamlRequests []docTypes.AIYAMLRequest, cfg docTypes.GeneratorConfig) error {
	if cfg.AIClient != nil && (len(codeRequests) > 0 || len(yamlRequests) > 0) {
		fmt.Printf("🚀 Sending %d code + %d YAML requests to AI\n", len(codeRequests), len(yamlRequests))
		processAIRequests(codeRequests, yamlRequests, cfg.AIClient)
		fmt.Println("✅ AI enhancement complete.")
	}
	return nil
}

func generateFinalDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig) error {
	fmt.Println("📝 Generating documentation output...")
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			lang := docUtils.GetLanguage(file.Path)
			if lang == "YAML" {
				fmt.Printf("📄 Generating YAML documentation for: %s\n", pkg.Name)
				if err := docGenerator.GenerateYAMLDoc(pkg, file.Path, cfg); err != nil {
					return fmt.Errorf("failed to generate YAML docs for package %s: %w", pkg.Name, err)
				}
			} else {
				fmt.Printf("📄 Generating code documentation for: %s\n", pkg.Name)
				if err := docGenerator.GeneratePackageDoc(pkg, file.Path, cfg); err != nil {
					return fmt.Errorf("failed to generate docs for package %s: %w", pkg.Name, err)
				}
			}
		}
	}
	fmt.Println("✅ All documentation successfully generated.")
	return nil
}
