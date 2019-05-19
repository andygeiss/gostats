package files

import "strings"

// FilterByExtension removes file names from a slice by a given file extension.
func FilterByExtension(ext string, in []string) (out []string) {
	for _, name := range in {
		if strings.HasSuffix(name, ".go") {
			out = append(out, name)
		}
	}
	return out
}
