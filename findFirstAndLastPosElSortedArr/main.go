package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{lowerBound(nums, target), upperBound(nums, target)}
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			ans = mid
			r = mid - 1
		}
	}

	return ans
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			ans = mid
			l = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(searchRange([]int{7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9}, 8))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
