package main

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	ans := 0
	for i, num := range nums1 {
		lo, hi := i, len(nums2)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if num <= nums2[mid] {
				ans = max(ans, mid-i)
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
}
