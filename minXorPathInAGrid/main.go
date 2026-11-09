package main

import (
	"fmt"
)

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]bool, m)
	for i := range m {
		dp[i] = make([][]bool, n)
		for j := range n {
			dp[i][j] = make([]bool, 1024)
		}
	}
	dp[0][0][grid[0][0]] = true
	for i := range m {
		for j := range n {
			for x := range 1024 {
				if !dp[i][j][x] {
					continue
				}

				if j+1 < n {
					xor := x ^ grid[i][j+1]
					dp[i][j+1][xor] = true
				}

				if i+1 < m {
					xor := x ^ grid[i+1][j]
					dp[i+1][j][xor] = true
				}
			}
		}
	}
	for x := range 1024 {
		if dp[m-1][n-1][x] {
			return x
		}
	}

	return -1
}

func main() {
	fmt.Println(minCost([][]int{{1, 2}, {3, 4}}))
}
