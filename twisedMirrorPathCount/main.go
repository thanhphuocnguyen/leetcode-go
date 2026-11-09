package main

import (
	"fmt"
)

const MOD = 1_000_000_007

func uniquePaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][2]int, m)
	for i := range m {
		dp[i] = make([][2]int, n)
		for j := range n {
			dp[i][j] = [2]int{-1, -1}

		}
	}
	return dfs(grid, dp, 0, 0, 0)
}

func dfs(grid [][]int, dp [][][2]int, r, c, dir int) int {
	if r >= len(grid) || c >= len(grid[0]) {
		return 0
	}

	if r == len(grid)-1 && c == len(grid[0])-1 {
		return 1
	}

	if dp[r][c][dir] != -1 {
		return dp[r][c][dir]
	}

	rs := 0
	if grid[r][c] == 0 {
		rs = dfs(grid, dp, r+1, c, 1) + dfs(grid, dp, r, c+1, 0)
	} else {
		if dir == 0 {
			rs = dfs(grid, dp, r+1, c, 1)
		} else {
			rs = dfs(grid, dp, r, c+1, 0)
		}
	}
	dp[r][c][dir] = rs % MOD

	return dp[r][c][dir]
}

func main() {
	fmt.Println(uniquePaths([][]int{{0, 1, 0}, {0, 0, 1}, {0, 0, 1}}))
}
