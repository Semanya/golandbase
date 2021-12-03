package analys

import (
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

func Analys(filepath string) (*AnalysResult, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	p := &AnalysResult{}
	ast.Inspect(node, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.GenDecl:
			if s.Tok == token.CONST || s.Tok == token.VAR {
				p.DeclCount++
			}
		case *ast.CallExpr:
			p.CallCount++
		case *ast.AssignStmt:
			p.AssignCount++
		case *ast.ImportSpec:
			p.ImportsCount++
		}
		return true
	})
	return p, nil
}
