package main

import "fmt"

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		suffixMin[n-i-1] = min(suffixMin[n-i], nums[n-i-1])
	}
	currMax := -1
	for i := range n {
		currMax = max(currMax, nums[i])
		if currMax-suffixMin[i] <= k {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstStableIndex([]int{5, 0, 1, 4}, 3))
}
