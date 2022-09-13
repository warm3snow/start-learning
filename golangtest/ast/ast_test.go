package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestGoAst(t *testing.T) {
	src := `
		package main
		func main() {
			c := a+b
		}
	`
	// Create the AST by parsing src.
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", src, 0)

	// Print the AST.
	ast.Print(fset, f)
}
