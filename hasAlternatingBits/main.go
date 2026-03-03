package main

import "fmt"

func hasAlternatingBits(n int) bool {
	ans := true
	prev := -1
	for n > 0 {
		if prev != -1 && prev == (n&1) {
			return false
		}
		prev = n & 1
		n >>= 1
	}
	return ans
}

func main() {
	fmt.Println(hasAlternatingBits(7))
}
