package main

import "fmt"

func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = 0
		}
	}

	// var solve func(l, r int) int
	// solve = func(l, r int) int {
	// 	if l > r {
	// 		return 0
	// 	}
	// 	if dp[l][r] != -1 {
	// 		return dp[l][r]
	// 	}
	// 	ans := 0
	// 	if s[l] != s[r] {
	// 		ans = 1 + min(solve(l+1, r), solve(l, r-1))
	// 	} else {
	// 		ans = solve(l+1, r-1)
	// 	}
	// 	dp[l][r] = ans
	// 	return ans
	// }
	// return solve(0, n-1)
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func main() {
	fmt.Println(minInsertions("mbadm"))
}
