package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	"github.com/MRGHOSJ/docupocus/internal/utils"
)

type GoAnalyzer struct{}

func (g *GoAnalyzer) Supports(projectDir string) bool {
	return exists(filepath.Join(projectDir, "go.mod")) || utils.HasFilesWithExtension(projectDir, ".go")
}

func (g *GoAnalyzer) Analyze(projectDir string) (*AnalyzerResult, error) {
	files, err := os.ReadDir(projectDir)
	if err != nil {
		return nil, err
	}

	result := &AnalyzerResult{}

	for _, file := range files {
		fullPath := filepath.Join(projectDir, file.Name())

		if file.IsDir() {
			// Recursively analyze subdirectories
			subResult, err := g.Analyze(fullPath)
			if err != nil {
				return nil, err
			}
			result.Merge(subResult)
			continue
		}

		if strings.HasSuffix(file.Name(), ".go") {
			fileResult, err := g.AnalyzeFile(fullPath)
			if err != nil {
				return nil, err
			}
			result.Files = append(result.Files, fileResult)
		}
	}

	return result, nil
}

// Merge merges another result into this one
func (r *AnalyzerResult) Merge(other *AnalyzerResult) {
	r.Files = append(r.Files, other.Files...)
}

// exists checks if a file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// AnalyzeFile analyzes a single Go file
func (g *GoAnalyzer) AnalyzeFile(path string) (*AnalyzedFile, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	pkg := Package{
		Name: filepath.Base(filepath.Dir(path)),
		Path: path,
	}

	// First pass: collect all function names
	funcNames := make(map[string]struct{})
	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			fullName := fn.Name.Name
			if fn.Recv != nil {
				// For methods, include receiver type
				recvType := utils.ExprToString(fn.Recv.List[0].Type)
				fullName = recvType + "." + fullName
			}
			funcNames[fullName] = struct{}{}
		}
	}

	// Second pass: analyze functions with call tracking
	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					typeSpec := spec.(*ast.TypeSpec)
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						s := Struct{
							Name: typeSpec.Name.Name,
							Doc:  ai.Documentation{Summary: utils.DocToString(d.Doc)},
						}
						// ... existing struct field processing ...
						pkg.Structs = append(pkg.Structs, s)
					}
				}
			}
		case *ast.FuncDecl:
			fn := Function{
				Name:       d.Name.Name,
				Doc:        ai.Documentation{Summary: utils.DocToString(d.Doc)},
				Receiver:   utils.RecvToString(d.Recv),
				Parameters: extractParams(d.Type.Params),
				Results:    extractParams(d.Type.Results),
				Calls:      []string{},
			}

			// Track function calls
			ast.Inspect(d.Body, func(n ast.Node) bool {
				if call, ok := n.(*ast.CallExpr); ok {
					if ident, ok := call.Fun.(*ast.Ident); ok {
						if _, exists := funcNames[ident.Name]; exists {
							fn.Calls = append(fn.Calls, ident.Name)
						}
					} else if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
						// Handle method calls
						if xIdent, ok := sel.X.(*ast.Ident); ok {
							fullName := xIdent.Name + "." + sel.Sel.Name
							if _, exists := funcNames[fullName]; exists {
								fn.Calls = append(fn.Calls, fullName)
							}
						}
					}
				}
				return true
			})

			pkg.Funcs = append(pkg.Funcs, fn)
		}
	}

	return &AnalyzedFile{
		Path:     path,
		Packages: []Package{pkg},
	}, nil
}

func extractParams(fl *ast.FieldList) []Parameter {
	var params []Parameter
	if fl == nil {
		return params
	}
	for _, f := range fl.List {
		paramType := utils.ExprToString(f.Type)
		if len(f.Names) == 0 {
			params = append(params, Parameter{
				Name: "",
				Type: paramType,
			})
		}
		for _, name := range f.Names {
			params = append(params, Parameter{
				Name: name.Name,
				Type: paramType,
			})
		}
	}
	return params
}
