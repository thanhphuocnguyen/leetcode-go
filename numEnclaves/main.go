package main

import "fmt"

var DIRECTIONS = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	for i, row := range grid {
		for j, cell := range row {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && cell == 1 && !visited[i][j] {
				dfs(grid, visited, m, n, i, j)
			}
		}
	}

	ans := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				ans++
			}
		}
	}

	return ans
}

func isValid(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

func dfs(grid [][]int, visited [][]bool, m, n, i, j int) {
	grid[i][j] = 0
	visited[i][j] = true
	for _, dir := range DIRECTIONS {
		ni, nj := i+dir[0], j+dir[1]

		if isValid(ni, nj, m, n) && !visited[ni][nj] && grid[ni][nj] == 1 {
			dfs(grid, visited, m, n, ni, nj)
		}
	}

}

func main() {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
}
