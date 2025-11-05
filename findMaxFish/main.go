package main

func findMaxFish(grid [][]int) int {
	numOfRows, numOfCols := len(grid), len(grid[0])
	visited := make([][]bool, numOfRows)
	for i := range visited {
		visited[i] = make([]bool, numOfCols)
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < numOfCols; j++ {
			if grid[i][j] != 0 {
				ans = max(ans, dfs(i, j, grid, visited))
			}
		}
	}
	return ans
}

func dfs(row, col int, grid [][]int, visited [][]bool) int {
	if !isValidCell(row, col, len(grid), len(grid[0])) || visited[row][col] || grid[row][col] == 0 {
		return 0
	}
	catches := grid[row][col]
	grid[row][col] = 0
	catches += dfs(row+1, col, grid, visited)
	catches += dfs(row-1, col, grid, visited)
	catches += dfs(row, col+1, grid, visited)
	catches += dfs(row, col-1, grid, visited)
	return catches
}

func isValidCell(row, col, numOfRows, numOfCols int) bool {
	return row >= 0 && col >= 0 && row < numOfRows && col < numOfCols
}

func main() {
	grid := [][]int{
		{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0},
	}
	println(findMaxFish(grid))
}
