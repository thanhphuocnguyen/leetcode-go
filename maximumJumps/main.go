package main

import (
	"fmt"
	"math"
)

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range n {
		memo[i] = -1
	}
	memo[0] = -1
	ans := dp(nums, memo, n, 0, target)
	if ans < 0 {
		return -1
	}
	return ans
}

func dp(nums, memo []int, n, i, target int) int {
	if i == len(nums)-1 {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}
	ans := math.MinInt32
	for j := i + 1; j < n; j++ {
		diff := nums[j] - nums[i] // nums[j] - nums[i]
		// take in rage
		if diff >= -target && diff <= target {
			// go to n to check
			ans = max(ans, 1+dp(nums, memo, n, j, target))
		}
	}
	memo[i] = ans
	return ans
}

func main() {
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 0))
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 2))
}
