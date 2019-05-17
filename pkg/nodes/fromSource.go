package nodes

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

	var state nodesVisitor
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			ast.Walk(&state, fn)
		}
	}

	return state.Nodes, nil
}

type nodesVisitor struct {
	Nodes int
}

func (c *nodesVisitor) Visit(node ast.Node) ast.Visitor {
	c.Nodes++
	return c
}
