package main

import "fmt"

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])
	r, c := click[0], click[1]

	if board[r][c] == 'M' {
		board[r][c] = 'X'
		return board
	}
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	dfs(board, visited, m, n, click[0], click[1])

	return board
}

func isValid(m, n, i, j int) bool {
	return i < m && j < n && i >= 0 && j >= 0
}

func dfs(board [][]byte, visited [][]bool, m, n, r, c int) {
	visited[r][c] = true
	mineCnt := 0
	for _, dir := range directions {
		nr, nc := r+dir[0], c+dir[1]
		if isValid(m, n, nr, nc) && !visited[nr][nc] {
			if board[nr][nc] == 'M' {
				mineCnt++
			}
		}
	}
	if mineCnt > 0 {
		board[r][c] = byte('0' + mineCnt)
	} else {
		board[r][c] = 'B'
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if isValid(m, n, nr, nc) && !visited[nr][nc] {
				dfs(board, visited, m, n, nr, nc)
			}
		}
	}
}

func main() {
	grid := updateBoard(
		[][]byte{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
		},
		[]int{3, 0})
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}
