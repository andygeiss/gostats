package files

import "strings"

// ExcludeByPrefix removes strings from the slice by a given prefix.
func ExcludeByPrefix(prefix string, in []string) (out []string) {
	for _, name := range in {
		if !strings.HasPrefix(name, prefix) {
			out = append(out, name)
		}
	}
	return out
}
