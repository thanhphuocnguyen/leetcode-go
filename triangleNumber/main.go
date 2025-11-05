package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			idx := binarySearch(nums, j+1, nums[i]+nums[j])
			ans += idx - j - 1
		}
	}
	return ans
}
func binarySearch(arr []int, left, target int) int {
	lo, hi := left, len(arr)-1
	res := len(arr)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// n1 + n2 > mid
		if arr[mid] >= target {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return res
}
func main() {
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}
