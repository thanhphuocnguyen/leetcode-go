package main

import (
	"fmt"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	n := len(nums)
	sort.Ints(nums)
	for original <= nums[n-1] && original >= nums[0] && bisect(nums, original) {
		original *= 2
	}
	return original
}

func bisect(nums []int, target int) bool {
	lo, hi := 0, len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if target == nums[mid] {
			return true
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
}
