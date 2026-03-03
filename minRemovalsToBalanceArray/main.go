package main

import (
	"fmt"
	"sort"
)

func minRemoval(nums []int, k int) int {
	n := len(nums)
	// sort nums
	sort.Ints(nums)
	ans := n - 1
	// sliding window to check the widest window that have k*min <= max
	l, r := 0, 0
	for r < n {
		for nums[r] > k*nums[l] {
			l++
		}
		// get min operation eq len of array remove size of window
		ans = min(ans, n-(r-l+1))
		r++
	}
	// return max answer
	return ans
}

func main() {
	fmt.Println(minRemoval([]int{2, 1, 5}, 2))
	fmt.Println(minRemoval([]int{1, 6, 2, 9}, 3))
}
