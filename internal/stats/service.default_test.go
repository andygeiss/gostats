package stats_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/gostats/internal/stats"
	"testing"
)

func TestDefaultService_Measure_Should_Return_Empty_Slice_If_No_Source_Found(t *testing.T) {
	service := stats.NewDefaultService("../../testdata/1")
	slice, err := service.Measure()
	assert.That(t, err, is.Nil())
	assert.That(t, slice, is.NotNil())
	assert.That(t, len(slice), is.Equal(0))
}

func TestDefaultService_Measure_Should_Return_Stats_For_One_Source_File(t *testing.T) {
	service := stats.NewDefaultService("../../testdata/3")
	slice, err := service.Measure()
	assert.That(t, err, is.Nil())
	assert.That(t, slice, is.NotNil())
	assert.That(t, len(slice), is.Equal(1))
	assert.That(t, slice[0].Ident, is.Equal("../../testdata/3/foo.go"))
}

func TestDefaultService_Measure_Should_Return_Complexity_Of_1_If_FuncDecl_Was_Found(t *testing.T) {
	service := stats.NewDefaultService("../../testdata/3")
	slice, err := service.Measure()
	assert.That(t, err, is.Nil())
	assert.That(t, slice, is.NotNil())
	assert.That(t, len(slice), is.Equal(1))
	assert.That(t, slice[0].Complexity, is.Equal(1))
}

func TestDefaultService_Measure_Should_Return_Nodes_10_If_FuncDecl_Was_Found(t *testing.T) {
	service := stats.NewDefaultService("../../testdata/3")
	slice, err := service.Measure()
	assert.That(t, err, is.Nil())
	assert.That(t, slice, is.NotNil())
	assert.That(t, len(slice), is.Equal(1))
	assert.That(t, slice[0].Nodes, is.Equal(10))
}
