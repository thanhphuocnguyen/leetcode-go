package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	n := len(nums)
	// memo will save the minimum step to reach to that index
	memo := make([]int, n)
	for i := range n {
		memo[i] = -1
	}

	return dp(nums, memo, 0)
}

func dp(nums []int, memo []int, idx int) int {
	if idx >= len(nums)-1 {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	// take
	take := math.MaxInt32
	for i := 1; i <= nums[idx]; i++ {
		take = min(take, 1+dp(nums, memo, idx+i))
	}
	// skip
	memo[idx] = take
	return memo[idx]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
