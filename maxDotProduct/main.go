package main

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1)+1, len(nums2)+1
	const MIN_VAL = -1_000_000_009
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = MIN_VAL
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			prod := nums1[i-1] * nums2[j-1]
			dp[i][j] = max(prod, prod+dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(maxDotProduct([]int{}, []int{}))
}
