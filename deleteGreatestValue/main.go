package main

import (
	"fmt"
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	for i := range m {
		sort.Ints(grid[i])
	}
	for j := range n {
		candidate := -1
		for i := range m {
			candidate = max(candidate, grid[i][j])
		}
		ans += candidate
	}

	return ans
}

func main() {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
}
