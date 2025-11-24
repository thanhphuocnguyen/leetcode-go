package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	memo := make(map[int]bool)
	return dp(nums, memo, n-1, 0)
}

func dp(nums []int, memo map[int]bool, last, currIdx int) bool {
	if currIdx >= last {
		return true
	}
	if memo[currIdx] {
		return memo[currIdx]
	}
	for i := 1; i <= nums[currIdx]; i++ {
		ans := dp(nums, memo, last, i+currIdx)
		if ans {
			memo[currIdx] = true
			return memo[currIdx]
		}
	}
	memo[currIdx] = false
	return false
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{1}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
