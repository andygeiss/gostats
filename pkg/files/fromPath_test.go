package files_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/gostats/pkg/files"
	"testing"
)

func TestFromPath_Should_Return_An_Empty_Slice_If_Dir_Is_1_And_Empty(t *testing.T) {
	slice, err := files.FromPath("../../testdata/1")
	assert.That(t, len(slice), is.Equal(0))
	assert.That(t, err, is.Nil())
}

func TestFromPath_Should_Return_Foo_If_Dir_Is_2(t *testing.T) {
	slice, err := files.FromPath("../../testdata/2")
	assert.That(t, len(slice), is.Equal(1))
	assert.That(t, slice[0], is.Equal("../../testdata/2/foo"))
	assert.That(t, err, is.Nil())
}
