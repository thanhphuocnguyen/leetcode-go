package main

import (
	"fmt"
	"math"
)

func maxSumDivThree(nums []int) int {
	n := len(nums)
	memo := make([][3]int, n)
	for i := range n {
		memo[i] = [3]int{-1, -1, -1}
	}
	return dp(nums, memo, 0, 0)
}

func dp(nums []int, memo [][3]int, i int, mod int) int {
	if i >= len(nums) {
		if mod == 0 {
			return 0
		}
		return -math.MinInt32
	}

	if memo[i][mod] != -1 {
		return memo[i][mod]
	}
	// take
	take := nums[i] + dp(nums, memo, i+1, (mod+nums[i])%3)
	//skip
	skip := dp(nums, memo, i+1, mod)
	memo[i][mod] = max(take, skip)
	return memo[i][mod]
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
}
