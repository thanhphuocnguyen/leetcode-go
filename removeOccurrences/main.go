package main

import "strings"

func removeOccurrences(s string, part string) string {
	idx := strings.Index(s, part)
	for idx >= 0 {
		s = s[:idx] + s[idx+len(part):]
		idx = strings.Index(s, part)
	}
	return s
}

func main() {
	removeOccurrences("daabcbaabcbc", "abc")
}
