package main

import "fmt"

var DIRS [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	ans := 0
	for i := range m {
		visited[i] = make([]bool, n)
	}
	for i := range m {
		for j := range n {
			if !visited[i][j] && grid[i][j] == 1 {
				ans = max(ans, dfs(visited, grid, i, j, m, n))
			}
		}
	}
	return ans
}

func isValid(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

func dfs(visited [][]bool, grid [][]int, i, j, m, n int) int {
	visited[i][j] = true
	area := 1
	for _, dir := range DIRS {
		r, c := dir[0], dir[1]
		ni, nj := i+r, j+c
		if isValid(ni, nj, m, n) && grid[ni][nj] == 1 && !visited[ni][nj] {
			area += dfs(visited, grid, ni, nj, m, n)
		}
	}
	return area
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}
