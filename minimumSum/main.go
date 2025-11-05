package main

import "math"

func minimumSum(nums []int) int {
	n := len(nums)
	prefixMin := make([]int, n)
	suffixMin := make([]int, n)
	prefixMin[0] = nums[0]
	suffixMin[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		prefixMin[i] = min(prefixMin[i-1], nums[i])
		suffixMin[n-1-i] = min(suffixMin[n-i], nums[n-1-i])
	}

	ans := math.MaxInt32
	for i := 1; i < n-1; i++ {
		if nums[i] > prefixMin[i-1] && nums[i] > suffixMin[i+1] {
			ans = min(ans, prefixMin[i-1]+nums[i]+suffixMin[i+1])
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	nums := []int{8, 6, 1, 5, 3}
	result := minimumSum(nums)
	println(result) // Output: 9
}
