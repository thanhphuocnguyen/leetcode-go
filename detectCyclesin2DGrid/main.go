package main

import "fmt"

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, cell := range row {
			visited := make([][]bool, m)
			for i1 := range m {
				visited[i1] = make([]bool, n)
			}
			if dfs(grid, visited, cell, -1, -1, i, j) {
				return true
			}
		}
	}

	return false
}
func isValid(m, n, i, j int) bool {
	return i < m && j < n && i >= 0 && j >= 0
}

func dfs(grid [][]byte, visited [][]bool, p byte, prevI, prevJ, i, j int) bool {
	visited[i][j] = true
	ans := false
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if isValid(len(grid), len(grid[0]), ni, nj) && (ni != prevI || nj != prevJ) && grid[ni][nj] == p {
			if visited[ni][nj] {
				return true

			}
			ans = ans || dfs(grid, visited, p, i, j, ni, nj)
			if ans {
				return true
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(containsCycle([][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}))
}
