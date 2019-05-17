package stats

import (
	"github.com/andygeiss/gostats/pkg/complexity"
	"github.com/andygeiss/gostats/pkg/files"
	"github.com/andygeiss/gostats/pkg/nodes"
)

type defaultService struct {
	path string
}

// Measure ...
func (s *defaultService) Measure() ([]Statistics, error) {
	allFiles, err := files.FromPath(s.path)
	if err != nil {
		return nil, err
	}

	result := files.FilterByExtension(".go", allFiles)
	result = files.ExcludeBySuffix("_test.go", result)
	result = files.ExcludeByPrefix("testdata", result)
	return createStatisticsBy(result), nil
}

// NewDefaultService ...
func NewDefaultService(path string) Service {
	return &defaultService{
		path: path,
	}
}

func createStatisticsBy(sourceFiles []string) []Statistics {
	var statistics []Statistics
	for _, file := range sourceFiles {
		c, err := complexity.FromSource(file)
		if err != nil {
			continue
		}
		n, err := nodes.FromSource(file)
		if err != nil {
			continue
		}
		statistics = append(statistics, Statistics{
			Ident:      file,
			Complexity: c,
			Nodes:      n,
		})
	}
	return statistics
}
