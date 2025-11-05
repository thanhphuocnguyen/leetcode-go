package main

import "sort"

func check(nums []int) bool {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	for offset := 0; offset < len(nums); offset++ {
		isSorted := true
		for i, sorted := range sortedNums {
			if nums[(i+offset)%len(nums)] != sorted {
				isSorted = false
				break
			}
		}
		if isSorted {
			return true
		}
	}
	return false
}
