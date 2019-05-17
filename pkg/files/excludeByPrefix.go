package files

import "strings"

// ExcludeByPrefix ...
func ExcludeByPrefix(prefix string, in []string) (out []string) {
	for _, name := range in {
		if !strings.HasPrefix(name, prefix) {
			out = append(out, name)
		}
	}
	return out
}
