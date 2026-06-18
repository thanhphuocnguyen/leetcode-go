package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid+1 < n && nums[mid+1] < nums[mid] {
			return nums[mid+1]
		} else if mid-1 >= 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		} else if nums[n-1] > nums[mid] {
			hi = mid - 1
		} else if nums[n-1] < nums[mid] {
			lo = mid + 1
		} else {
			return nums[mid]
		}
	}
	return nums[lo]
}

func main() {
	fmt.Println(findMin([]int{1}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
