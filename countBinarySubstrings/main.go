package main

import "fmt"

func countBinarySubstrings(s string) int {
	ones, zeros := 0, 0
	ans := 0
	if s[0] == '0' {
		zeros = 1
	} else {
		ones = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if s[i] == '0' {
				zeros++
			} else {
				ones++
			}
		} else {
			ans += min(ones, zeros)
			if s[i] == '0' {
				zeros = 1
			} else {
				ones = 1
			}
		}
	}
	ans += min(ones, zeros)
	return ans
}

func main() {
	fmt.Println(countBinarySubstrings("10101"))
}
