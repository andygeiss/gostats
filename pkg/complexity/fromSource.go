package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// FromSource parses a file by a given path and returns the complexity.
func FromSource(path string) (int, error) {
	f, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return 0, err
	}
	var state complexityVisitor
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			ast.Walk(&state, fn)
		}
	}
	return state.Complexity, nil
}

type complexityVisitor struct {
	Complexity int
}

// Visit handles each node in the abstract syntax tree and increases the complexity for each
// * CaseClause
// * CommClause
// * ForStmt
// * FuncDecl
// * IfStmt
// * RangeStmt
// * SwitchStmt
// * TypeSwitchStmt
// * BinaryExpr with AND or OR
func (c *complexityVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.CaseClause, *ast.CommClause, *ast.ForStmt, *ast.FuncDecl, *ast.IfStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt:
		c.Complexity++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			c.Complexity++
		}
	}
	return c
}
