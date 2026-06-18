package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	lt, gt := 0, len(nums)-1
	ans := make([]int, len(nums))
	eqCnt := 0
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] < pivot {
			ans[lt] = nums[i]
			lt++
		}
		if nums[j] > pivot {
			ans[gt] = nums[j]
			gt--
		}
		if nums[i] == pivot {
			eqCnt++
		}
	}
	for i := 0; i < eqCnt; i++ {
		ans[lt] = pivot
		lt++
	}
	return ans
}

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
}
