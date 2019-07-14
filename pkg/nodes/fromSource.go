package nodes

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// FromSource parses all files and calculates the number of nodes by a given path.
func FromSource(path string) (int, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, path, nil, 0)
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

// Visit counts the number of each node visited.
func (c *nodesVisitor) Visit(node ast.Node) ast.Visitor {
	c.Nodes++
	return c
}
