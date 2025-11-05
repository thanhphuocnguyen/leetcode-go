package main

func solveNQueens(n int) [][]string {
	board := make([][]bool, n)
	for i := range n {
		board[i] = make([]bool, n)
	}
	isSafe := func(row, col int, board [][]bool) bool {
		n := len(board)
		i, j := 0, 0
		// check upper pos at column
		for i = 0; i < row; i++ {
			if board[i][col] {
				return false
			}
		}

		// check upper left diagonal
		for i, j = row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] {
				return false
			}
		}

		// check upper right diagonal
		for i, j = row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] {
				return false
			}
		}
		return true
	}
	ans := make([][]string, 0)
	var backtrack func(row, n int, board [][]bool)
	backtrack = func(row, n int, board [][]bool) {
		if row == n {
			validBoard := make([]string, n)
			for i, row := range board {
				bs := make([]byte, n)
				for j, cell := range row {
					if cell {
						bs[j] = 'Q'
					} else {
						bs[j] = '.'
					}
				}
				validBoard[i] = string(bs)
			}
			ans = append(ans, validBoard)
			return
		}
		for col := 0; col < n; col++ {
			if isSafe(row, col, board) {
				board[row][col] = true
				backtrack(row+1, n, board)
				// backtrack
				board[row][col] = false
			}
		}
	}
	backtrack(0, n, board)

	return ans
}

func main() {
	solveNQueens(4)
}
