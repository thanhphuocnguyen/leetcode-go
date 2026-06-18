package main

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := range m {
		dp[i] = make([][]int, n)
		for j := range n {
			dp[i][j] = make([]int, k+1)
			for c := range k + 1 {
				dp[i][j][c] = -1
			}
		}
	}

	dp[0][0][0] = 0

	for i := range m {
		for j := range n {
			for c := range k + 1 {
				if dp[i][j][c] == -1 {
					continue
				}

				// move down
				if i+1 < m {
					val := grid[i+1][j]
					cost := 0
					if val != 0 {
						cost = 1
					}
					nextCost := cost + c
					if nextCost <= k {
						dp[i+1][j][nextCost] = max(dp[i+1][j][nextCost], dp[i][j][c]+val)
					}
				}

				// move right
				if j+1 < n {
					val := grid[i][j+1]
					cost := 0
					if val != 0 {
						cost = 1
					}
					nextCost := cost + c
					if nextCost <= k {
						dp[i][j+1][nextCost] = max(dp[i][j+1][nextCost], dp[i][j][c]+val)
					}
				}
			}
		}
	}
	ans := -1
	for c := range k + 1 {
		ans = max(ans, dp[m-1][n-1][c])
	}
	return ans
}

func main() {

}
