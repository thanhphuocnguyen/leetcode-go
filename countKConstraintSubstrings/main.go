package main

import "fmt"

func countKConstraintSubstrings(s string, k int) int {
	start, end := 0, 0
	ones, zeros := 0, 0
	ans := 0
	for end < len(s) {
		if s[end] == '1' {
			ones++
		} else {
			zeros++
		}
		for ones > k && zeros > k {
			if s[start] == '1' {
				ones--
			} else {
				zeros--
			}
			start++
		}
		ans += end - start + 1
		end++
	}
	return ans
}

func main() {
	fmt.Println(countKConstraintSubstrings("1010101", 2))
	fmt.Println(countKConstraintSubstrings("10101", 1))
}
