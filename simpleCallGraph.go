package main

import {
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
}

func main() {
	srcPath := os.Args[1]

	// Parse the program into an AST
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, srcPath, nil, 0)
	if err != nil {
		panic(err)
	}

	// Loop over all the declarations in the program
	for _, f := range node.Decls {
		// Check if current declaration is a function
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// If yes, print the name
		fmt.Printf(fn.Name.Name + "->")
	}
	fmt.Fprint(os.Stdout, "\033[2D")
}