package main

const MOD int = 1_000_000_007

func countPartitions(nums []int, k int) int {
	memo := make([]int, len(nums))
	return dp(memo, len(nums), 0, k)
}

func dp(memo []int, n, idx, k int) int {
	if idx == n {

	}
	return -1
}
