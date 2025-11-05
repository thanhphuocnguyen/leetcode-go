package main

import "fmt"

func pivotInteger(n int) int {
	prefixSum := 0
	for i := 1; i <= n; i++ {
		prefixSum += i
	}

	suffixSum := 0
	for i := n; i >= 1; i-- {
		prefixSum -= i
		if prefixSum == suffixSum {
			return i
		}
		suffixSum += i
	}

	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
}
