package main

import (
	"fmt"
	"math"
)

func minOperations(nums []int) int {
	n := len(nums)
	ones := 0
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}
	// if there are ones in array then just return the min ops
	// is the len of arr - ones count, bc every gcd of num with 1 is one
	if ones > 0 {
		return n - ones
	}

	// other just find the shortest sub array that have gcd = 1 by brute force
	ans := math.MaxInt
	for i, num := range nums {
		g := num
		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				ans = min(ans, (j-i)+n-1)
				break
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	fmt.Println(minOperations([]int{2, 10, 6, 14}))
	fmt.Println(minOperations([]int{2, 6, 3, 4}))
}
