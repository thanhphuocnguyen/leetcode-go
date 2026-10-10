package main

import (
	"fmt"
	"math"
)

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := range n {
		dp[i] = math.MaxInt32
	}

	prefixSum := 0
	ans := n + 1
	for i, j := 0, 0; j < n; j++ {
		prefixSum += arr[j]
		for prefixSum > target {
			prefixSum -= arr[i]
			i++
		}
		dp[j+1] = dp[j]
		if prefixSum == target {
			dist := j - i + 1
			ans = min(ans, dist+dp[i])
			dp[j+1] = min(dist, dp[j])
		}
	}

	if ans == n+1 {
		return -1
	}

	return ans
}

func main() {
	fmt.Println(minSumOfLengths([]int{3, 1, 1, 1, 5, 1, 2, 1}, 3))
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))
}
