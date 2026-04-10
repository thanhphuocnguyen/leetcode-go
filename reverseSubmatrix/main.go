package main

import "fmt"

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := 0; i < k/2; i++ {
		for j := y; j < y+k; j++ {
			grid[x+i][j], grid[(x+k-1)-i][j] = grid[(x+k-1)-i][j], grid[x+i][j]
		}
	}
	return grid
}

func main() {
	fmt.Println(reverseSubmatrix([][]int{{6, 16, 14}, {1, 2, 19}, {14, 17, 15}, {18, 7, 6}, {14, 12, 5}}, 2, 1, 2))
	fmt.Println(reverseSubmatrix([][]int{{3, 4, 2, 3}, {2, 3, 4, 2}}, 0, 2, 2))
	fmt.Println(reverseSubmatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 1, 0, 3))
}
