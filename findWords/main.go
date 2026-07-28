package main

import (
	"fmt"
	"unicode"
)

func findWords(words []string) []string {
	mp := map[rune]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}
	ans := make([]string, 0)
	for _, word := range words {
		currRow := -1
		valid := true
		for _, c := range word {
			lowerC := unicode.ToLower(c)
			if currRow == -1 {
				currRow = mp[lowerC]
			} else if currRow != mp[lowerC] {
				valid = false
				break
			}
		}
		if valid {
			ans = append(ans, word)
		}
	}
	return ans
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
