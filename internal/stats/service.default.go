package stats

import (
	"github.com/andygeiss/goinfo/pkg/complexity"
	"github.com/andygeiss/goinfo/pkg/files"
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
	return createStatisticsBy(files.FilterBy(".go", allFiles)), nil
}

// NewDefaultService ...
func NewDefaultService(path string) Service {
	return &defaultService{
		path:path,
	}
}

func createStatisticsBy(sourceFiles []string) []Statistics {
	var statistics []Statistics
	for _, file := range sourceFiles {
		comp, err := complexity.FromSource(file)
		if err != nil {
			continue
		}
		statistics = append(statistics, Statistics{
			Ident:      file,
			Complexity: comp,
		})
	}
	return statistics
}
