package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	n := len(nums)
	memo := make([]int, n)

	return dp(nums, memo, n, 0)
}

func dp(nums, memo []int, n, i int) int {
	if i == n-1 {
		return 1
	}
	if memo[i] != 0 {
		return memo[i]
	}
	ans := 1
	for j := i + 1; j < n; j++ {
		if nums[j] > nums[i] {
			// take
			ans = max(ans, 1+dp(nums, memo, n, j))
		} else {
			// skip
			ans = max(ans, dp(nums, memo, n, j))
		}
	}
	memo[i] = ans
	return memo[i]
}
func main() {
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS([]int{0}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
