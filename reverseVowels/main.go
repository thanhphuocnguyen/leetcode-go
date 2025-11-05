package main

import "fmt"

func reverseVowels(s string) string {
	if len(s) < 2 {
		return s
	}
	chars := make([]byte, len(s))

	i, j := 0, len(s)-1
	for i <= j {
		chars[i] = s[i]
		chars[j] = s[j]

		if !isVowel(s[i]) {
			chars[i] = s[i]
			i++
		}
		if !isVowel(s[j]) {
			chars[j] = s[j]
			j--
		}
		if isVowel(s[i]) && isVowel(s[j]) {
			chars[i] = s[j]
			chars[j] = s[i]
			i++
			j--
		}
	}

	return string(chars)
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

func isVowel(c byte) bool {
	c = toLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	fmt.Println(reverseVowels("1b1"))
	fmt.Println(reverseVowels("a a"))
	fmt.Println(reverseVowels("a,"))
	fmt.Println(reverseVowels(".,"))
	fmt.Println(reverseVowels(" "))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("IceCreAm"))
}
