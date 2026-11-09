package main

import (
	"container/list"
	"fmt"
)

var DIR = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	for i := range m {
		if !visited[i][0] && grid[i][0] == 0 {
			bfs(grid, visited, i, 0)
		}

		if !visited[i][n-1] && grid[i][n-1] == 0 {
			bfs(grid, visited, i, n-1)
		}
	}
	for j := range n {
		if !visited[0][j] && grid[0][j] == 0 {
			bfs(grid, visited, 0, j)
		}

		if !visited[m-1][j] && grid[m-1][j] == 0 {
			bfs(grid, visited, m-1, j)
		}
	}

	// exclude all island from the border
	ans := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !visited[i][j] && grid[i][j] == 0 {
				bfs(grid, visited, i, j)
				ans++
			}
		}

	}

	return ans
}

func bfs(grid [][]int, visited [][]bool, i, j int) {
	q := list.New()
	q.PushBack([2]int{i, j})
	for q.Len() > 0 {
		el := q.Front()
		pos := el.Value.([2]int)
		q.Remove(el)
		if visited[pos[0]][pos[1]] {
			continue
		}

		visited[pos[0]][pos[1]] = true

		for _, dir := range DIR {
			nextRow, nextCol := pos[0]+dir[0], pos[1]+dir[1]
			if nextRow > 0 && nextCol > 0 && nextRow < len(grid) && nextCol < len(grid[0]) && !visited[nextRow][nextCol] && grid[nextRow][nextCol] == 0 {
				q.PushBack([2]int{nextRow, nextCol})
			}
		}
	}
}

func main() {
	fmt.Println(closedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}))
}
