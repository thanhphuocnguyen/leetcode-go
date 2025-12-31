package main

import "fmt"

func countNegatives(grid [][]int) int {
	ans := 0
	for _, nums := range grid {
		ans += len(nums) - binarySearch(nums, -1)
	}
	return ans
}

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
}
