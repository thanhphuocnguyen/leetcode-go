package main

import "fmt"

func countKeyChanges(s string) int {
	prev := toLower(s[0])
	ans := 0
	for i := 1; i < len(s); i++ {
		curr := toLower(s[i])
		if curr != prev {
			ans++
			prev = curr
		}
	}
	return ans
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' - 'A' + b
	}
	return b
}
func main() {
	fmt.Println(countKeyChanges("aAbBcC"))
}
