package main

func swimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	minElevation, maxElevation := 0, 0
	for i := range n {
		visited[i] = make([]bool, n)
		for j := range n {
			maxElevation = max(maxElevation, grid[i][j])
		}
	}
	ans := -1
	for minElevation <= maxElevation {
		avgElevation := minElevation + (maxElevation-minElevation)/2
		if hasPath(grid, visited, n, 0, 0, avgElevation) {
			ans = avgElevation
			maxElevation = avgElevation - 1
		} else {
			minElevation = avgElevation + 1
		}
	}
	return ans
}

func hasPath(grid [][]int, visited [][]bool, n, row, col, waterElevation int) bool {
	if row < 0 || col < 0 || row >= n || col >= n || visited[row][col] || grid[row][col] > waterElevation {
		return false
	}
	if row == n-1 && col == n-1 {
		return true
	}
	if hasPath(grid, visited, n, row-1, col, waterElevation) ||
		hasPath(grid, visited, n, row+1, col, waterElevation) ||
		hasPath(grid, visited, n, row, col-1, waterElevation) ||
		hasPath(grid, visited, n, row, col+1, waterElevation) {
		return true
	}
	return false
}

func main() {
	swimInWater([][]int{{0, 2}, {1, 3}})
}
