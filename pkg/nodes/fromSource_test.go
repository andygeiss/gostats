package nodes_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/gostats/pkg/nodes"
	"testing"
)

func TestFromSource_Should_Return_10_If_FuncDecl_Was_Found(t *testing.T) {
	result, err := nodes.FromSource("../../testdata/1/foo.go")
	assert.That(t, result, is.Equal(10))
	assert.That(t, err, is.Nil())
}