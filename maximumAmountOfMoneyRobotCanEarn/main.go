package main

func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	// dp[i][j][k] represents the maximum amount of coins we can collect from (0, 0) to (i, j) with k skips
	dp := make([][][]int, m)
	for i := range m {
		dp[i] = make([][]int, n)
		for j := range n {
			dp[i][j] = make([]int, 3)
			for k := range 3 {
				dp[i][j][k] = -1 << 60
			}
		}
	}
	dp[0][0][0] = max(coins[0][0], 0)
	// skip the first cell
	dp[0][0][1] = dp[0][0][0]
	dp[0][0][2] = dp[0][0][1]
	// fill the dp table
	for i := range m {
		for j := range n {
			for k := range 3 {
				if i > 0 {
					// take the current cell
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k]+coins[i][j])
					// skip the current cell
					if k > 0 {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k-1])
					}
				}
				if j > 0 {
					// take the current cell
					dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k]+coins[i][j])
					// skip the current cell
					if k > 0 {
						dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k-1])
					}
				}
			}
		}
	}
	// the answer is the maximum amount of coins we can collect at (m-1, n-1) with 0, 1, or 2 skips
	last := dp[m-1][n-1]

	return max(last[0], last[1], last[2])
}
