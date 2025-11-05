package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	// sort array
	sort.Ints(nums)
	// init ans is always 1 element
	ans := 0
	// using binary search to find lower bound and upper bound of current element
	for i, num := range nums {
		lb := lowerBound(nums, i, num-k)
		ub := upperBound(nums, i, num+k)
		start, end := i, i
		if lb != len(nums) {
			start = lb
		}
		if ub != len(nums) {
			end = ub
		}
		curr := 1
		for i := start; i < end; i++ {
			if nums[i] == nums[i+1] {
				curr++
			} else if numOperations > 0 {
				curr++
				numOperations--
			}
		}
		ans = max(ans, curr)
	}

	return ans
}

func lowerBound(nums []int, end, target int) int {
	lo, hi := 0, end
	ans := len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func upperBound(nums []int, start, target int) int {
	ans := len(nums)
	lo, hi := start, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] <= target {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFrequency([]int{2}, 7, 0))
	fmt.Println(maxFrequency([]int{5, 11, 20, 20}, 5, 1))
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2))
}
