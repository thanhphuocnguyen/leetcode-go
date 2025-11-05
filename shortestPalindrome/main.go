package main

import "strings"

func shortestPalindrome(s string) string {
	reversed := Reverse(s)
	combined := s + "#" + reversed
	longestCount := kmp(combined)
	var sb strings.Builder
	sb.WriteString(reversed)
	return reversed[:len(s)-longestCount] + s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func kmp(s string) int {
	prefixTable := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		length := prefixTable[i-1]
		println(string(s[i]), string(s[length]))
		for length > 0 && s[i] != s[length] {
			length = prefixTable[length-1]
		}
		if s[i] == s[length] {
			length++
		}
		prefixTable[i] = length
	}

	return prefixTable[len(s)-1]
}

func main() {
	shortestPalindrome("aacecaaa")
	shortestPalindrome("abcd")
}
