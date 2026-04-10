package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := range m {
		for j := range n {
			if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}
			if j > 0 && i == 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
