package main

import (
	"fmt"
)

func maxRotateFunction(nums []int) int {
	n := len(nums)
	// f := func(k int) int {
	// 	ans := 0
	// 	start := n - k
	// 	for i := range n {
	// 		ans += nums[(start+i)%n] * i
	// 	}

	// 	return ans
	// }
	sum := 0
	f := 0
	for i, num := range nums {
		sum += num
		f += num * i
	}
	ans := f
	for i := 1; i < n; i++ {
		f = f + sum - n*nums[n-i]
		ans = max(ans, f)
	}

	return ans
}

func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
}
