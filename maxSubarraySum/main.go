package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
	// approach prefix sum and sliding window & kadane's algorithm
	n := len(nums)
	minPrefixSum := make([]int64, k)
	for i := range k {
		minPrefixSum[i] = math.MaxInt / 2
	}
	minPrefixSum[k-1] = 0
	var prefixSum int64
	var ans int64 = math.MinInt64
	for i := 0; i < n; i++ {
		num := int64(nums[i])
		prefixSum += num
		ans = max(ans, prefixSum-minPrefixSum[i%k])
		minPrefixSum[i%k] = min(minPrefixSum[i%k], prefixSum)
	}

	return ans
}

func main() {
	fmt.Println(maxSubarraySum([]int{-1, -2, -3, -4, -5}, 4))
	fmt.Println(maxSubarraySum([]int{-5, 1, 2, -3, 4}, 2))
}
