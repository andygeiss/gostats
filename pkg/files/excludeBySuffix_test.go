package files_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/gostats/pkg/files"
	"testing"
)

func TestExcludeBySuffix_Should_Return_An_Empty_Slice_If_Input_Was_Empty(t *testing.T) {
	out := files.ExcludeBySuffix("", []string{})
	assert.That(t, len(out), is.Equal(0))
}

func TestExcludeBySuffix_Should_Return_One_Go_File_If_Input_Excludes_A_Test_File(t *testing.T) {
	out := files.ExcludeBySuffix("_test.go", []string{"foo.go", "foo_test.go"})
	assert.That(t, len(out), is.Equal(1))
	assert.That(t, out[0], is.Equal("foo.go"))
}
