package main

import "fmt"

func minCapability(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		minVal = min(num, minVal)
		maxVal = max(num, maxVal)
	}

	for minVal < maxVal {
		mid := minVal + (maxVal-minVal)/2
		if check(nums, k, mid) {
			maxVal = mid
		} else {
			minVal = mid + 1
		}
	}
	return minVal
}
func check(nums []int, k int, guess int) bool {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		cnt += (nums[i] - nums[i-1] - 1) / guess
	}
	return cnt >= k
}

func main() {
	fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
}
