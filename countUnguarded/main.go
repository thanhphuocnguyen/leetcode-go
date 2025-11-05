package main

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, w := range walls {
		grid[w[0]][w[1]] = 2
	}
	ans := n*m - len(walls)
	for _, g := range guards {
		row, col := g[0], g[1]
		for i := row; i < m; i++ {
			if grid[i][col] == 2 {
				break
			}
			if grid[i][col] == 1 {
				continue
			}
			grid[i][col] = 1
			ans--
		}
		for i := row; i >= 0; i-- {
			if grid[i][col] == 2 {
				break
			}
			if grid[i][col] == 1 {
				continue
			}
			grid[i][col] = 1
			ans--
		}
		for i := col; i < n; i++ {
			if grid[row][i] == 2 {
				break
			}
			if grid[row][i] == 1 {
				continue
			}
			grid[row][i] = 1
			ans--
		}
		for i := col; i >= 0; i-- {
			if grid[row][i] == 2 {
				break
			}
			if grid[row][i] == 1 {
				continue
			}
			grid[row][i] = 1
			ans--
		}
	}
	return ans
}

func main() {
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}
