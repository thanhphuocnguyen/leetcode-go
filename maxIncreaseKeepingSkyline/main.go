package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n, ans := len(grid), 0
	maxInRow := make([]int, n)
	maxInCol := make([]int, n)
	for i, row := range grid {
		for j, cell := range row {
			maxInRow[i] = max(maxInRow[i], cell)
			maxInCol[j] = max(maxInCol[j], cell)
		}
	}

	for i := range n {
		for j := range n {
			ans += min(maxInRow[i], maxInCol[j]) - grid[i][j]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}
