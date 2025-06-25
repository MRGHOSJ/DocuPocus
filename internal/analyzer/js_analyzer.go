package analyzer

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/utils"
)

type JSAnalyzer struct{}

func (j *JSAnalyzer) Supports(projectDir string) bool {
	return utils.HasFilesWithExtension(projectDir, ".js") || utils.HasFilesWithExtension(projectDir, ".ts")
}

func (j *JSAnalyzer) Analyze(projectDir string) (*AnalyzerResult, error) {
	var result AnalyzerResult

	err := filepath.WalkDir(projectDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !(strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".ts")) {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		funcs := extractJSFunctions(string(content))
		structs := extractJSClasses(string(content))

		result.Files = append(result.Files, &AnalyzedFile{
			Path: path,
			Packages: []Package{
				{
					Name:    filepath.Base(filepath.Dir(path)),
					Path:    path,
					Structs: structs,
					Funcs:   funcs,
				},
			},
		})

		return nil
	})

	return &result, err
}

// -- Helpers --

func extractJSFunctions(src string) []Function {
	funcRegex := regexp.MustCompile(`(?m)^(?:function\s+(\w+)|(?:const|let|var)\s+(\w+)\s*=\s*function)\s*\(([^)]*)\)`)
	jsdocRegex := regexp.MustCompile(`(?s)/\*\*(.*?)\*/`)
	callRegex := regexp.MustCompile(`(?m)(\b\w+\b)\s*\(`)

	// First pass: collect all function names
	funcNames := make(map[string]struct{})
	for _, match := range funcRegex.FindAllStringSubmatch(src, -1) {
		name := match[1]
		if name == "" {
			name = match[2]
		}
		funcNames[name] = struct{}{}
	}

	// Second pass: analyze functions with call tracking
	var funcs []Function
	for _, match := range funcRegex.FindAllStringSubmatch(src, -1) {
		name := match[1]
		if name == "" {
			name = match[2]
		}
		rawParams := match[3]
		params := parseParamList(name, rawParams)
		doc := findDocBefore(match[0], src, jsdocRegex)

		// Find all calls within this function's body
		var calls []string
		funcBody := extractJSFunctionBody(src, match[0])
		for _, callMatch := range callRegex.FindAllStringSubmatch(funcBody, -1) {
			callee := callMatch[1]
			if _, exists := funcNames[callee]; exists && callee != name {
				calls = append(calls, callee)
			}
		}

		funcs = append(funcs, Function{
			Name:       name,
			Parameters: params,
			Doc:        ai.Documentation{Summary: doc},
			Calls:      calls,
		})
	}
	return funcs
}

func extractJSFunctionBody(src, funcDecl string) string {
	// Find the function declaration
	declIndex := strings.Index(src, funcDecl)
	if declIndex == -1 {
		return ""
	}

	// Find the start of the function body (after the parameters)
	bodyStart := declIndex + len(funcDecl)
	for bodyStart < len(src) && src[bodyStart] != '{' {
		bodyStart++
	}
	if bodyStart >= len(src) {
		return ""
	}
	bodyStart++ // Skip the '{'

	// Extract the body by tracking braces
	var body strings.Builder
	braceLevel := 1
	for i := bodyStart; i < len(src) && braceLevel > 0; i++ {
		if src[i] == '{' {
			braceLevel++
		} else if src[i] == '}' {
			braceLevel--
		}
		if braceLevel > 0 {
			body.WriteByte(src[i])
		}
	}

	return body.String()
}

func extractJSClasses(src string) []Struct {
	classRegex := regexp.MustCompile(`(?m)^class\s+(\w+)`)
	jsdocRegex := regexp.MustCompile(`(?s)/\*\*(.*?)\*/`)
	fieldRegex := regexp.MustCompile(`(?m)^\s*(?:static\s+)?(\w+)\s*(?:=\s*[^;]+)?;`)
	methodRegex := regexp.MustCompile(`(?m)^\s*(?:static\s+)?(\w+)\s*\(([^)]*)\)\s*{`)

	var structs []Struct
	for _, match := range classRegex.FindAllStringSubmatch(src, -1) {
		className := match[1]
		doc := findDocBefore(match[0], src, jsdocRegex)

		s := Struct{
			Name:    className,
			Doc:     ai.Documentation{Summary: doc},
			Fields:  []Field{},
			Methods: []Function{},
		}

		// Find the class body
		classBody := extractJsClassBody(src, match[0])

		// Extract fields
		for _, fieldMatch := range fieldRegex.FindAllStringSubmatch(classBody, -1) {
			fieldName := fieldMatch[1]
			s.Fields = append(s.Fields, Field{
				Name: fieldName,
				Type: "unknown",
			})
		}

		// Extract methods
		for _, methodMatch := range methodRegex.FindAllStringSubmatch(classBody, -1) {
			methodName := methodMatch[1]
			rawParams := methodMatch[2]
			params := parseParamList(methodName, rawParams)

			methodDoc := findDocBefore(methodMatch[0], classBody, jsdocRegex)

			s.Methods = append(s.Methods, Function{
				Name:       methodName,
				Parameters: params,
				Doc:        ai.Documentation{Summary: methodDoc},
			})
		}

		structs = append(structs, s)
	}
	return structs
}

func extractJsClassBody(src, classDecl string) string {
	// Find the class declaration
	declIndex := strings.Index(src, classDecl)
	if declIndex == -1 {
		return ""
	}

	// Find the start of the class body (after the class name)
	bodyStart := declIndex + len(classDecl)
	for bodyStart < len(src) && src[bodyStart] != '{' {
		bodyStart++
	}
	if bodyStart >= len(src) {
		return ""
	}
	bodyStart++ // Skip the '{'

	// Extract the body by tracking braces
	var body strings.Builder
	braceLevel := 1
	for i := bodyStart; i < len(src) && braceLevel > 0; i++ {
		if src[i] == '{' {
			braceLevel++
		} else if src[i] == '}' {
			braceLevel--
		}
		if braceLevel > 0 {
			body.WriteByte(src[i])
		}
	}

	return body.String()
}

func parseParamList(_ string, raw string) []Parameter {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	var params []Parameter
	for _, p := range parts {
		p = strings.TrimSpace(p)
		params = append(params, Parameter{Name: p, Type: "unknown"})
	}
	return params
}

func findDocBefore(context, src string, docRegex *regexp.Regexp) string {
	index := strings.Index(src, context)
	if index == -1 {
		return ""
	}
	before := src[:index]
	if matches := docRegex.FindAllStringSubmatch(before, -1); len(matches) > 0 {
		return strings.TrimSpace(matches[len(matches)-1][1])
	}
	return ""
}
