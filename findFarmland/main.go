package main

import "fmt"

func findFarmland(land [][]int) [][]int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(land), len(land[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	isValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	var dfs func(i, j int, coord []int)

	dfs = func(i, j int, coord []int) {
		visited[i][j] = true
		for _, dir := range directions {
			nextRow, nextCol := i+dir[0], j+dir[1]
			if isValid(nextRow, nextCol) && !visited[nextRow][nextCol] && land[nextRow][nextCol] != 0 {
				coord[2] = max(nextRow, coord[2])
				coord[3] = max(nextCol, coord[3])
				dfs(nextRow, nextCol, coord)
			}
		}

	}
	ans := make([][]int, 0)
	for i := range m {
		for j := range n {
			if land[i][j] == 1 && !visited[i][j] {
				coord := []int{i, j, i, j}
				dfs(i, j, coord)
				ans = append(ans, coord)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findFarmland([][]int{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}}))
	fmt.Println(findFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
}
