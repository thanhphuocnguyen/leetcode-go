package main

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	totalSum := 0
	for _, r := range s1 {
		totalSum += int(r)
	}
	for _, r := range s2 {
		totalSum += int(r)
	}
	m, n := len(s1)+1, len(s2)+1
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = int(s1[i-1]) + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return totalSum - 2*dp[m-1][n-1]
}
