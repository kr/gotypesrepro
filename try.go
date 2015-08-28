// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	name := "testdata/import.go"

	var files []*ast.File
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, name, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	files = append(files, file)

	config := types.Config{
		Importer: importer.Default(),
		Error:    func(error) {},
	}
	info := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
	}
	pkg, err := config.Check("testdata", fs, files, info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pkg)
}
