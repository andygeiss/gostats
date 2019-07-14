package files_test

import (
	"github.com/andygeiss/golib/pkg/assert"
	"github.com/andygeiss/gostats/pkg/files"
	"testing"
)

func TestFromPath_Should_Return_An_Empty_Slice_If_Dir_Is_1_And_Empty(t *testing.T) {
	slice, err := files.FromPath("../../testdata/1")
	assert.That(t, err, nil)
	assert.That(t, len(slice), 0)
}

func TestFromPath_Should_Return_Foo_If_Dir_Is_2(t *testing.T) {
	slice, err := files.FromPath("../../testdata/2")
	assert.That(t, err, nil)
	assert.That(t, len(slice), 1)
	assert.That(t, slice[0], "../../testdata/2/foo")
}
