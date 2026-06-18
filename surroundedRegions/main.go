package main

import "fmt"

var DIRECTIONS = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	for i := range m {
		if !visited[i][0] && board[i][0] == 'O' {
			dfs(board, visited, m, n, i, 0)
		}
		if !visited[i][n-1] && board[i][n-1] == 'O' {
			dfs(board, visited, m, n, i, n-1)
		}
	}
	for j := range n {
		if !visited[0][j] && board[0][j] == 'O' {
			dfs(board, visited, m, n, 0, j)
		}
		if !visited[m-1][j] && board[m-1][j] == 'O' {
			dfs(board, visited, m, n, m-1, j)
		}
	}
	for i, row := range visited {
		for j, v := range row {
			if !v {
				board[i][j] = 'X'
			}
		}
	}

}

func isValid(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

func dfs(board [][]byte, visited [][]bool, m, n, i, j int) {
	visited[i][j] = true
	for _, dir := range DIRECTIONS {
		ni, nj := i+dir[0], j+dir[1]
		if isValid(ni, nj, m, n) && !visited[ni][nj] && board[ni][nj] == 'O' {
			dfs(board, visited, m, n, ni, nj)
		}
	}
}

func main() {
	grid := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(grid)
	fmt.Println(grid)
}
