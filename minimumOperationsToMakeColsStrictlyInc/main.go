package main

import "fmt"

func minimumOperations(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ops := 0
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			if grid[i][j] <= grid[i-1][j] {
				ops += (grid[i-1][j] - grid[i][j]) + 1
				grid[i][j] = grid[i-1][j] + 1
			}
		}
	}

	return ops
}

func main() {
	fmt.Println(minimumOperations([][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}))
}
