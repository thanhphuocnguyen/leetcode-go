package main

func pivotArray(nums []int, pivot int) []int {
	lt, gt := 0, len(nums)-1
	ans := make([]int, len(nums))
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j+1 {
		if nums[i] < pivot {
			ans[lt] = nums[i]
			lt++
		} else if nums[j] > pivot {
			ans[gt] = nums[j]
			gt--
		}
	}

	return ans
}
