package main

import "fmt"

var DIRECTIONS = [][][]int{
	{{0, -1}, {0, 1}},  // 1 left right
	{{1, 0}, {-1, 0}},  // 2 down up
	{{0, -1}, {1, 0}},  // 3 left down
	{{1, 0}, {0, 1}},   // 4 down right
	{{0, -1}, {-1, 0}}, // 5 left up
	{{-1, 0}, {0, 1}},  // 6 up right
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	return dfs(grid, visited, m, n, 0, 0)
}

func isValid(m, n, i, j int) bool {
	return i < m && j < n && i >= 0 && j >= 0
}

func dfs(grid [][]int, visited [][]bool, m, n, i, j int) bool {
	if i == m-1 && j == n-1 {
		return true
	}
	visited[i][j] = true
	for _, dirs := range DIRECTIONS[grid[i][j]-1] {
		ni, nj := i+dirs[0], j+dirs[1]
		if !isValid(m, n, ni, nj) || visited[ni][nj] {
			continue
		}
		for _, dirs2 := range DIRECTIONS[grid[ni][nj]-1] {
			// one of paths of next cell must connect with current cell
			ni2, nj2 := ni+dirs2[0], nj+dirs2[1]
			if ni2 == i && nj2 == j {
				if dfs(grid, visited, m, n, ni, nj) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(hasValidPath([][]int{{3, 4, 3, 4}, {2, 2, 2, 2}, {6, 5, 6, 5}}))
	fmt.Println(hasValidPath([][]int{{2}, {2}, {2}, {2}, {2}, {2}, {6}}))
	fmt.Println(hasValidPath([][]int{{1, 1, 1, 1, 1, 1, 3}}))
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}}))
}
