package main

import (
	"fmt"
	"sort"
)

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	diagonalMap := make(map[int][]int)
	for i, row := range grid {
		for j, col := range row {
			diagonalMap[i-j] = append(diagonalMap[i-j], col)
		}
	}

	diagonalIndexMap := make(map[int]int)

	for k, v := range diagonalMap {
		diagonalIndexMap[k] = 0
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}

	for i := range n {
		for j := range n {
			key := i - j
			grid[i][j] = diagonalMap[key][diagonalIndexMap[key]]
			diagonalIndexMap[key]++
		}
	}
	return grid
}

func main() {
	fmt.Println(sortMatrix([][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}))
}
