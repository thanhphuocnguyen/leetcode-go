package main

import "strings"

func repeatLimitedString(s string, repeatLimit int) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}

	var sb strings.Builder
	curIdx := 'z' - 'a'
	for curIdx >= 0 {
		if freq[curIdx] == 0 {
			curIdx--
			continue
		}
		take := min(freq[curIdx], repeatLimit)
		freq[curIdx] -= take
		for i := 0; i < take; i++ {
			sb.WriteRune('a' + curIdx)
		}
		if take < repeatLimit || freq[curIdx] == 0 {
			curIdx--
			continue
		}
		next := curIdx - 1
		for next >= 0 && freq[next] == 0 {
			next--
		}
		if next < 0 {
			break
		}
		sb.WriteRune('a' + next)
		freq[next]--
	}
	return sb.String()
}
