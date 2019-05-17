package files_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/gostats/pkg/files"
	"testing"
)

func TestFilterByExtension_Should_Return_An_Empty_Slice_If_Input_Was_Empty(t *testing.T) {
	out := files.FilterByExtension("", []string{})
	assert.That(t, len(out), is.Equal(0))
}

func TestFilterByExtension_Should_Return_One_Go_File_If_Input_Has_One_Go_File(t *testing.T) {
	out := files.FilterByExtension(".go", []string{"foo.go"})
	assert.That(t, len(out), is.Equal(1))
	assert.That(t, out[0], is.Equal("foo.go"))
}

func TestFilterByExtension_Should_Return_One_Go_File_If_Input_Has_One_Go_File_And_One_Without_Go_Ext(t *testing.T) {
	out := files.FilterByExtension(".go", []string{"foo", "foo.go"})
	assert.That(t, len(out), is.Equal(1))
	assert.That(t, out[0], is.Equal("foo.go"))
}

func TestFilterByExtension_Should_Return_Two_Go_File_If_Input_Has_Two_Go_File_And_One_Without_Go_Ext(t *testing.T) {
	out := files.FilterByExtension(".go", []string{"foo", "foo.go", "bar.go"})
	assert.That(t, len(out), is.Equal(2))
	assert.That(t, out[0], is.Equal("foo.go"))
	assert.That(t, out[1], is.Equal("bar.go"))
}
