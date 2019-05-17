package files

import "strings"

// ExcludeBySuffix ...
func ExcludeBySuffix(suffix string, in []string) (out []string) {
	for _, name := range in {
		if !strings.HasSuffix(name, suffix) {
			out = append(out, name)
		}
	}
	return out
}
