package main

import (
	"fmt"
	"math"
)

func countTriples(n int) int {
	ans := 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if isPerfectSquare(float64(i*i) + float64(j*j)) {
				ans++
			}
		}
	}

	return ans
}
func isPerfectSquare(n float64) bool {
	if n < 0 {
		return false
	}
	sqrt := math.Sqrt(n)
	intRoot := int(sqrt)
	return intRoot*intRoot == int(n)
}

func main() {
	fmt.Println(countTriples(10))
	fmt.Println(countTriples(5))
}
