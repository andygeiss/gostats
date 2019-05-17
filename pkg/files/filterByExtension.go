package files

import "strings"

// FilterBy ...
func FilterBy(ext string, in []string) (out []string) {
	for _, name := range in {
		if strings.HasSuffix(name, ".go") {
			out = append(out, name)
		}
	}
	return out
}
