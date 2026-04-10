package main

import "fmt"

func countSubmatrices(grid [][]int, k int) int {
	// pre-compute prefix sum on grid to avoid re-calculate each time check sub-matrices
	m, n := len(grid), len(grid[0])
	ans := 0

	for i := range m {
		for j := range n {
			if i > 0 {
				grid[i][j] += grid[i-1][j]
			}

			if j > 0 {
				grid[i][j] += grid[i][j-1]
			}

			if i > 0 && j > 0 {
				grid[i][j] -= grid[i-1][j-1]
			}
			if grid[i][j] <= k {
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(countSubmatrices([][]int{{7, 6, 3}, {6, 6, 1}}, 18))
	fmt.Println(countSubmatrices([][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20))
}
