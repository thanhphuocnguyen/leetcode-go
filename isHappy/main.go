package main

import "fmt"

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	mp := make(map[int]bool)
	for !mp[n] {
		mp[n] = true
		sum := 0
		for n > 0 {
			remainder := n % 10
			n /= 10
			sum += remainder * remainder
		}
		if sum == 1 {
			return true
		}
		n = sum
	}
	return false
}

func main() {
	fmt.Println(isHappy(19))
}
