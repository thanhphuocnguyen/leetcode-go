package main

import "fmt"

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = -1
		}
	}
	return solve(dp, s, t, 0, 0)
}

func solve(dp [][]int, s, t string, i, j int) int {
	if j == len(t) {
		return 1
	}
	if i == len(s) {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}
	ans := 0
	if s[i] != t[j] {
		ans += solve(dp, s, t, i+1, j)
	} else {
		// take
		ans += solve(dp, s, t, i+1, j+1)
		// skip
		ans += solve(dp, s, t, i+1, j)
	}
	dp[i][j] = ans
	return ans
}

func main() {
	fmt.Println(numDistinct("babgbag", "bag"))
}
