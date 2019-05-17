package complexity_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/goinfo/pkg/complexity"
	"testing"
)

func TestFromSource_Should_Return_Complexity_1_If_One_FuncDecl_Was_Found(t *testing.T) {
	comp, err := complexity.FromSource("../../testdata/3/foo.go")
	assert.That(t, err, is.Nil())
	assert.That(t, comp, is.Equal(1))
}

func TestFromSource_Should_Return_Complexity_2_If_Two_FuncDecl_Was_Found(t *testing.T) {
	comp, err := complexity.FromSource("../../testdata/4/foo.go")
	assert.That(t, err, is.Nil())
	assert.That(t, comp, is.Equal(2))
}

func TestFromSource_Should_Return_Complexity_1_If_FuncDecl_With_SimpleStmt_Was_Found(t *testing.T) {
	comp, err := complexity.FromSource("../../testdata/5/foo.go")
	assert.That(t, err, is.Nil())
	assert.That(t, comp, is.Equal(1))
}

func TestFromSource_Should_Return_Complexity_2_If_FuncDecl_With_IfStmt_Was_Found(t *testing.T) {
	comp, err := complexity.FromSource("../../testdata/6/foo.go")
	assert.That(t, err, is.Nil())
	assert.That(t, comp, is.Equal(2))
}
