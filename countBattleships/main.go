package main

import "fmt"

var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func countBattleships(board [][]byte) int {
	// count battle ship by row or col:
	/* there are two battle ship at (0,0) and (3, [0,1,2])
	X . . . X
	. . . . X
	. . . . X
	*/
	ans := 0
	for i, row := range board {
		for j, cell := range row {
			if cell == 'X' {
				ans++
				dfs(board, i, j)
			}
		}
	}

	return ans
}

func dfs(board [][]byte, row, col int) {
	// stop when see a space
	if board[row][col] == '.' {
		return
	}
	// toggle visited cell
	board[row][col] = '.'
	// traversal by horizontal -

	for _, d := range dirs {
		nextRow, nextCol := row+d[0], col+d[1]
		if isValid(nextRow, nextCol, len(board), len(board[0])) && board[nextRow][nextCol] == 'X' {
			dfs(board, nextRow, nextCol)
		}
	}
}

func isValid(row, col, m, n int) bool {
	return row >= 0 && col >= 0 && row < m && col < n
}

func main() {
	fmt.Println(countBattleships([][]byte{{'.'}}))
	fmt.Println(countBattleships([][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}))
}
