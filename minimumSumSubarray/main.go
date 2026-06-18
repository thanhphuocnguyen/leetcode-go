package main

import (
	"fmt"
	"math"
)

func minimumSumSubarray(nums []int, l int, r int) int {
	// Calculate the prefix sum of the array
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = nums[i-1] + prefixSum[i-1]
	}

	ans := math.MaxInt32
	// Sliding window size l -> r
	for sz := l; sz <= r; sz++ {
		for i := range n - sz + 1 {
			sum := prefixSum[i+sz] - prefixSum[i]
			if sum > 0 {
				ans = min(ans, sum)
			}
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSumSubarray([]int{3, -2, 1, 4}, 2, 3))
}
