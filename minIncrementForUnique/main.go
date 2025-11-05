package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	prev := nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= prev {
			ans += prev + 1 - nums[i]
			nums[i] = prev + 1
		}
		prev = nums[i]
	}
	return ans
}

func main() {
	println(minIncrementForUnique([]int{1, 2, 2}))
}
