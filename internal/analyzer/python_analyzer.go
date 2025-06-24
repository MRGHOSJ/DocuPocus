package analyzer

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/utils"
)

type PythonAnalyzer struct{}

func (p *PythonAnalyzer) Supports(projectDir string) bool {
	return utils.HasFilesWithExtension(projectDir, ".py")
}

func (p *PythonAnalyzer) Analyze(projectDir string) (*AnalyzerResult, error) {
	var result AnalyzerResult

	err := filepath.WalkDir(projectDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".py") {
			return nil
		}

		fileContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		classes := extractPythonClasses(string(fileContent))
		funcs := extractPythonFunctions(string(fileContent))

		result.Files = append(result.Files, &AnalyzedFile{
			Path: path,
			Packages: []Package{
				{
					Name:    filepath.Base(filepath.Dir(path)),
					Path:    path,
					Structs: classes,
					Funcs:   funcs,
				},
			},
		})

		return nil
	})

	return &result, err
}

// -- Helpers --

func extractPythonClasses(src string) []Struct {
	classRegex := regexp.MustCompile(`(?m)^class\s+(\w+)\s*(\(.*\))?:\s*(?:#.*)?`)
	docstringRegex := regexp.MustCompile(`(?m)(?:"""|''')(.*?)("""|''')`)

	var structs []Struct
	for _, match := range classRegex.FindAllStringSubmatch(src, -1) {
		s := Struct{
			Name: match[1],
			Doc:  ai.Documentation{Summary: findDocstringAfter(match[0], src, docstringRegex)},
		}
		structs = append(structs, s)
	}
	return structs
}

func extractPythonFunctions(src string) []Function {
	funcRegex := regexp.MustCompile(`(?m)^def\s+(\w+)\s*\(([^)]*)\):`)
	docstringRegex := regexp.MustCompile(`(?m)(?:"""|''')(.*?)("""|''')`)
	callRegex := regexp.MustCompile(`(?m)(\b\w+\b)\s*\(`)

	// First pass: collect all function names
	funcNames := make(map[string]struct{})
	for _, match := range funcRegex.FindAllStringSubmatch(src, -1) {
		funcNames[match[1]] = struct{}{}
	}

	// Second pass: analyze functions with call tracking
	var funcs []Function
	for _, match := range funcRegex.FindAllStringSubmatch(src, -1) {
		name := match[1]
		rawParams := match[2]
		params := parseParamList(name, rawParams)

		// Find all calls within this function's body
		var calls []string
		funcBody := extractFunctionBody(src, match[0])
		for _, callMatch := range callRegex.FindAllStringSubmatch(funcBody, -1) {
			callee := callMatch[1]
			if _, exists := funcNames[callee]; exists && callee != name {
				calls = append(calls, callee)
			}
		}

		f := Function{
			Name:       name,
			Parameters: params,
			Doc:        ai.Documentation{Summary: findDocstringAfter(match[0], src, docstringRegex)},
			Calls:      calls,
		}
		funcs = append(funcs, f)
	}
	return funcs
}

func extractFunctionBody(src, funcDecl string) string {
	// Find the function declaration
	declIndex := strings.Index(src, funcDecl)
	if declIndex == -1 {
		return ""
	}

	// Find the start of the function body (after the colon)
	bodyStart := declIndex + len(funcDecl)
	for bodyStart < len(src) && (src[bodyStart] == ' ' || src[bodyStart] == '\n') {
		bodyStart++
	}

	// Find the indentation level of the function
	indentLevel := 0
	for i := declIndex - 1; i >= 0; i-- {
		if src[i] == '\n' {
			break
		}
		indentLevel++
	}

	// Extract the body by tracking indentation
	var body strings.Builder
	currentIndent := 0
	inBody := false

	for i := bodyStart; i < len(src); i++ {
		if src[i] == '\n' {
			currentIndent = 0
			for i+1 < len(src) && src[i+1] == ' ' {
				currentIndent++
				i++
			}
			if !inBody {
				if currentIndent > indentLevel {
					inBody = true
				} else {
					break
				}
			} else if currentIndent <= indentLevel {
				break
			}
		}
		if inBody {
			body.WriteByte(src[i])
		}
	}

	return body.String()
}

func findDocstringAfter(context, src string, docRegex *regexp.Regexp) string {
	index := strings.Index(src, context)
	if index == -1 {
		return ""
	}
	after := src[index+len(context):]
	if match := docRegex.FindStringSubmatch(after); match != nil {
		return strings.TrimSpace(match[1])
	}
	return ""
}
