package files

import (
	"os"
	"path/filepath"
)

// FromPath collects a slice of file names by a given path.
func FromPath(path string) (slice []string, err error) {
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		slice = append(slice, path)
		return nil
	}); err != nil {
		return nil, err
	}
	return
}
