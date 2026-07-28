package main

import "fmt"

func removeDuplicateLetters(s string) string {
	last := make([]int, 26)
	for i := range 26 {
		last[i] = -1
	}
	for i, r := range s {
		last[r-'a'] = i
	}
	seen := make([]bool, 26)
	st := make([]rune, 0)
	for i, r := range s {
		if seen[r-'a'] {
			continue
		}

		n := len(st)
		for n > 0 && last[st[n-1]-'a'] != -1 && r < st[n-1] && last[st[n-1]-'a'] > i {
			seen[st[n-1]-'a'] = false
			st = st[:n-1]
			n--
		}
		st = append(st, r)
		seen[r-'a'] = true
	}

	return string(st)
}

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
