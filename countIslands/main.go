package main

import "fmt"

func countIslands(grid [][]int, k int) int {
	var DIR = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])
	ans := 0
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int) int64

	dfs = func(i, j int) int64 {
		if visited[i][j] {
			return 0
		}
		visited[i][j] = true
		ans := int64(grid[i][j])
		for _, dir := range DIR {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && !visited[ni][nj] && grid[ni][nj] != 0 {
				ans += dfs(ni, nj)
			}
		}
		return ans
	}

	for i := range m {
		for j := range n {
			if !visited[i][j] && grid[i][j] != 0 {
				if dfs(i, j)%int64(k) == 0 {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countIslands([][]int{{0, 2, 1, 0, 0}, {0, 5, 0, 0, 5}, {0, 0, 1, 0, 0}, {0, 1, 4, 7, 0}, {0, 2, 0, 0, 8}}, 5))
}
