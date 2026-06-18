package main

import "fmt"

func greatestLetter(s string) string {
	mp := make(map[rune]bool)
	ans := 'A' - 1
	offset := 'a' - 'A'
	for _, rn := range s {
		if !isUpper(rn) {
			if _, ok := mp[rn-offset]; ok {
				ans = max(ans, rn)
			}
		} else {
			if _, ok := mp[rn+offset]; ok {
				ans = max(ans, rn+offset)
			}
		}
		mp[rn] = true
	}
	if ans == 'A'-1 {
		return ""
	}
	return string(ans - offset)
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func main() {
	fmt.Println(greatestLetter("arRAzFif"))
	fmt.Println(greatestLetter("lEeTcOdE"))
	fmt.Println(greatestLetter("AbCdEfGhIjK"))
}
