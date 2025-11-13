package main

import "fmt"

func maxOperations(s string) int {
	// loop from left to right
	// count the adjacent 1 first
	// if there are s[i] == 1 and s[i+1] == 0 so add the number of moves
	// continue skip 0 til we meet new 1
	n, ones, ans, i := len(s), 0, 0, 0
	for i < n {
		for i < n && s[i] == '1' {
			ones++
			i++
		}
		if i < n {
			ans += ones
		}
		for i < n && s[i] == '0' {
			i++
		}
	}
	return ans
}

func main() {
	fmt.Println(maxOperations("00111"))
	fmt.Println(maxOperations("1001101"))
}
