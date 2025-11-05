package main

import "fmt"

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := make([][]bool, len(grid))
	for i := range m {
		visited[i] = make([]bool, n)
	}
	ans := 0
	isValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != -1
	}
	var backtrack func(i, j, cntNeed, cnt int)
	backtrack = func(i, j, cntNeed, cnt int) {
		if grid[i][j] == 2 {
			if cnt == cntNeed {
				ans++
			}
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		for _, dir := range directions {
			nextRow, nextCol := i+dir[0], j+dir[1]
			if isValid(nextRow, nextCol) {
				backtrack(nextRow, nextCol, cntNeed, cnt+1)
			}
		}
		visited[i][j] = false
	}
	startRow, startCol, cntNeed := -1, -1, 2
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				startRow, startCol = i, j
			}
			if grid[i][j] == 0 {
				cntNeed++
			}
		}
	}
	backtrack(startRow, startCol, cntNeed, 1)
	return ans
}

func main() {
	fmt.Println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}))
}
