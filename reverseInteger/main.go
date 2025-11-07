package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}
	ans := 0
	for x > 0 {
		ans = ans*10 + x%10
		if ans > math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	if isNeg {
		return -ans
	}
	return ans
}

func main() {
	fmt.Println(reverse(1563847412))
}
