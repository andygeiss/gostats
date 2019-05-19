package files

import "strings"

// ExcludeBySuffix removes strings from a slice by a given suffix.
func ExcludeBySuffix(suffix string, in []string) (out []string) {
	for _, name := range in {
		if !strings.HasSuffix(name, suffix) {
			out = append(out, name)
		}
	}
	return out
}
