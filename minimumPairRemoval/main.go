package main

import (
	"slices"
	"sort"
)

func minimumPairRemoval(nums []int) int {
	ans := 0

	for len(nums) > 1 && !sort.IntsAreSorted(nums) {
		sum := nums[0] + nums[1]
		idx := 0
		for i := 1; i < len(nums)-1; i++ {
			if nums[i]+nums[i+1] < sum {
				sum = nums[i] + nums[i+1]
				idx = i
			}
		}
		nums[idx] = sum
		nums = slices.Delete(nums, idx+1, idx+2)
		ans++
	}

	return ans
}

func main() {
	minimumPairRemoval([]int{2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1})
	minimumPairRemoval([]int{5, 2, 3, 1})
}
