package main

import "fmt"

func projectionArea(grid [][]int) int {
	n := len(grid)
	ans := 0
	for i := range n {
		maxRow := 0
		maxCol := 0
		for j := range n {
			if grid[i][j] != 0 {
				ans++
			}
			maxCol = max(maxCol, grid[i][j])
			maxRow = max(maxRow, grid[j][i])

		}
		ans += maxCol
		ans += maxRow
	}
	return ans
}

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}}))
	fmt.Println(projectionArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
}
