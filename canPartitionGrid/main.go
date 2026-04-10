package main

import "fmt"

func canPartitionGrid(grid [][]int) bool {
	// calculate prefix sum for whole grid
	// then start splitting each cut to check if there any eq parts.
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
			if j > 0 {
				grid[i][j] += grid[i][j-1]
			}
			if j > 0 && i > 0 {
				grid[i][j] -= grid[i-1][j-1]
			}
		}
	}

	for i := 0; i < m; i++ {
		topPart := grid[i][n-1]
		bottomPart := grid[m-1][n-1] - topPart
		if topPart == bottomPart {
			return true
		}
	}

	for j := 0; j < n; j++ {
		leftPart := grid[m-1][j]
		rightPart := grid[m-1][n-1] - leftPart
		if rightPart == leftPart {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canPartitionGrid([][]int{{1, 4}, {2, 3}}))
}
