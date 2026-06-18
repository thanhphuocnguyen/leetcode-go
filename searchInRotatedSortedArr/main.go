package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[l] <= nums[m] {
			// left half is sorted
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				// search from right half
				l = m + 1
			}
		} else {
			// right half is sorted
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				// search from left half
				r = m - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{6, 7, 1, 2, 3, 4, 5}, 6))
	fmt.Println(search([]int{1, 3}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{5, 1, 3}, 5))
}
