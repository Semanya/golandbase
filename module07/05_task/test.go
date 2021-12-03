package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type AnalysResult struct {
	DeclCount    int
	CallCount    int
	AssignCount  int
	ImportsCount int
}

func main() {
	Analys("/Users/u19502010/rebrainme/golandbase/module07/02_task/assets/astanalys/astanalys.go")
}

func Analys(filepath string) (*AnalysResult, error) {
	mmm, zzz, sss, ggg := 0, 0, 0, 0
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.GenDecl:
			if s.Tok == token.CONST || s.Tok == token.VAR {
				sss += 1
			}
			fmt.Println("Decl")
			fmt.Println(s)
		case *ast.CallExpr:
			ggg += 1
			fmt.Println("call")
			fmt.Println(s)
		case *ast.AssignStmt:
			mmm += 1
			fmt.Println("assign")
			fmt.Println(s)
		case *ast.ImportSpec:
			zzz += 1
			fmt.Println("assign")
			fmt.Println(s)
		default:
			return true
		}
		return true
	})
	p := &AnalysResult{
		DeclCount:    sss,
		CallCount:    ggg,
		AssignCount:  mmm,
		ImportsCount: zzz,
	}

	fmt.Println("Call ", ggg)
	fmt.Println("Decl ", sss)
	fmt.Println("Imports", zzz)
	fmt.Println("Assign", mmm)
	return p, nil
}
