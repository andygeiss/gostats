package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// FromSource ...
func FromSource(path string) (int, error) {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
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

func (c *complexityVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.FuncDecl, *ast.IfStmt:
		c.Complexity++
	}
	return c
}
