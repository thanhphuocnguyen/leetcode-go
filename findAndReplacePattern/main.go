package main

import "fmt"

var s1MapS2 = make(map[rune]rune)
var s2MapS1 = make(map[rune]rune)

func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0)
	for _, word := range words {
		if isIsomorphic(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func isIsomorphic(s1, s2 string) bool {
	clear(s1MapS2)
	clear(s2MapS1)
	for i, r1 := range s1 {
		r2 := rune(s2[i])
		r1MapR2, ok := s1MapS2[r1]
		if !ok {
			s1MapS2[r1] = r2
			r1MapR2 = r2
		}
		r2MapR1, ok := s2MapS1[r2]
		if !ok {
			s2MapS1[r2] = r1
			r2MapR1 = r1
		}
		if r1MapR2 != r2 || r2MapR1 != r1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}
