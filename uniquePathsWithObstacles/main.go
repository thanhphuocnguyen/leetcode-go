package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := range m {
		memo[i] = make([]int, n)
		for j := range n {
			memo[i][j] = -1
		}
	}

	return dp(obstacleGrid, memo, 0, 0, m, n)
}

func dp(grid [][]int, memo [][]int, i, j, m, n int) int {
	if i >= m || j >= n {
		return 0
	}
	if i == m-1 && j == n-1 && grid[i][j] == 0 {
		return 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if grid[i][j] == 1 {
		return 0
	}
	down := dp(grid, memo, i+1, j, m, n)
	right := dp(grid, memo, i, j+1, m, n)
	memo[i][j] = down + right
	return memo[i][j]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
