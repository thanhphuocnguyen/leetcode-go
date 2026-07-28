package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k = k % (m * n)
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
	}
	for i := range m {
		for j := range n {
			nextCol := j + k
			nextRow := i
			if nextCol >= n {
				nextRow = (nextRow + (nextCol / n)) % m
				nextCol = nextCol % n
			}
			ans[nextRow][nextCol%n] = grid[i][j]
		}
	}
	return ans
}

func main() {
	fmt.Println(shiftGrid([][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}, 23))
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
}
