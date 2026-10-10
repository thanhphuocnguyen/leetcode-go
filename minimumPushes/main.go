package main

import "fmt"

func minimumPushes(word string) int {
	ans := 0
	n := len(word) / 8
	rem := len(word) % 8
	ans = (n*(n+1)/2)*8 + (rem * (n + 1))
	return ans
}

func main() {
	fmt.Println(minimumPushes("xycdefghij"))
}
