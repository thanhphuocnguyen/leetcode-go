package main

import (
	"fmt"
)

func maxProductPath(grid [][]int) int {
	const MOD = 1_000_000_007
	m, n := len(grid), len(grid[0])
	dpMin := make([][]int64, m)
	dpMax := make([][]int64, m)
	for i := range m {
		dpMin[i] = make([]int64, n)
		dpMax[i] = make([]int64, n)
	}
	dpMax[0][0] = int64(grid[0][0])
	dpMin[0][0] = int64(grid[0][0])
	for i := 1; i < m; i++ {
		dpMax[i][0] = (dpMax[i-1][0] * int64(grid[i][0]))
		dpMin[i][0] = dpMax[i][0]
	}

	for j := 1; j < n; j++ {
		dpMax[0][j] = (dpMax[0][j-1] * int64(grid[0][j]))
		dpMin[0][j] = dpMax[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] >= 0 {
				maxPrev := max(dpMax[i-1][j], dpMax[i][j-1])
				minPrev := min(dpMin[i-1][j], dpMin[i][j-1])
				dpMax[i][j] = (maxPrev * int64(grid[i][j]))
				dpMin[i][j] = (minPrev * int64(grid[i][j]))
			} else {
				maxPrev := max(dpMax[i-1][j], dpMax[i][j-1])
				minPrev := min(dpMin[i-1][j], dpMin[i][j-1])
				dpMax[i][j] = (minPrev * int64(grid[i][j]))
				dpMin[i][j] = (maxPrev * int64(grid[i][j]))
			}
		}
	}
	if dpMax[m-1][n-1] < 0 {
		return -1
	}

	return int(dpMax[m-1][n-1] % MOD)
}

func main() {
	fmt.Println(maxProductPath([][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}))
	fmt.Println(maxProductPath([][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}))
}
