package files

import "strings"

// FilterByExtension ...
func FilterByExtension(ext string, in []string) (out []string) {
	for _, name := range in {
		if strings.HasSuffix(name, ".go") {
			out = append(out, name)
		}
	}
	return out
}
