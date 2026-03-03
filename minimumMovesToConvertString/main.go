package main

import "fmt"

func minimumMoves(s string) int {
	ans := 0
	i := 0
	for i < len(s) {
		if s[i] == 'X' {
			ans++
			i += 3
		} else {
			i++
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumMoves("OOOO"))
	fmt.Println(minimumMoves("XXOX"))
	fmt.Println(minimumMoves("XXX"))
}
