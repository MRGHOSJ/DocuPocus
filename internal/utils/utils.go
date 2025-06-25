package utils

import (
	"go/ast"
	"os"
	"path/filepath"
	"strings"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func DocToString(doc *ast.CommentGroup) string {
	if doc != nil {
		return strings.TrimSpace(doc.Text())
	}
	return ""
}

func TagToString(tag *ast.BasicLit) string {
	if tag != nil {
		return tag.Value
	}
	return ""
}

func ExprToString(expr ast.Expr) string {
	if expr == nil {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(ExprToRaw(expr), "\n", " "))
}

func ExprToRaw(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return ExprToString(e.X) + "." + e.Sel.Name
	case *ast.StarExpr:
		return "*" + ExprToString(e.X)
	case *ast.ArrayType:
		return "[]" + ExprToString(e.Elt)
	case *ast.MapType:
		return "map[" + ExprToString(e.Key) + "]" + ExprToString(e.Value)
	case *ast.FuncType:
		return "func"
	default:
		return ""
	}
}

func RecvToString(recv *ast.FieldList) string {
	if recv == nil || len(recv.List) == 0 {
		return ""
	}
	return ExprToString(recv.List[0].Type)
}

func HasFilesWithExtension(dir, ext string) bool {
	found := false
	filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(d.Name(), ext) {
			found = true
			return filepath.SkipDir
		}
		return nil
	})
	return found
}

func FieldTagToString(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	// Remove backticks and quotes
	tagValue := strings.Trim(tag.Value, "`\"")
	return tagValue
}
