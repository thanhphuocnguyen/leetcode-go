package main

import "fmt"

func solveSudoku(board [][]byte) {
	n := len(board)
	solved := false
	seenRow := make([][]bool, n)
	seenCol := make([][]bool, n)
	seenBox := make([][]bool, n)
	for i := range n {
		seenRow[i] = make([]bool, n+1)
		seenCol[i] = make([]bool, n+1)
		seenBox[i] = make([]bool, n+1)
	}
	getBoxIdx := func(row, col int) int {
		return (row/n)*n + (col / n)
	}
	for i := range n {
		for j := range n {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				seenRow[i][num] = true
				seenCol[j][num] = true
				seenBox[getBoxIdx(i, j)][num] = true
			}
		}
	}

	canPlace := func(row, col, num int) bool {
		boxIdx := getBoxIdx(row, col)
		return !seenRow[row][num] && !seenCol[col][num] && !seenBox[boxIdx][num]
	}

	placeValue := func(row, col int, num int) {
		board[row][col] = byte('0' + num)
		boxIdx := getBoxIdx(row, col)
		seenRow[row][num] = true
		seenCol[col][num] = true
		seenBox[boxIdx][num] = true
	}

	unplaceValue := func(row, col int, num int) {
		boxIdx := getBoxIdx(row, col)
		seenRow[row][num] = false
		seenCol[col][num] = false
		seenBox[boxIdx][num] = false
		board[row][col] = '.'
	}
	var backtrack func(row, col int)
	var placeNextCell func(row, col int)

	backtrack = func(row, col int) {
		if board[row][col] == '.' {
			for i := 1; i <= 9; i++ {
				if canPlace(row, col, i) {
					placeValue(row, col, i)
					placeNextCell(row, col)
					if !solved {
						unplaceValue(row, col, i)
					}
				}
			}
		} else {
			placeNextCell(row, col)
		}
	}

	placeNextCell = func(row, col int) {
		if row == n-1 && col == n-1 {
			solved = true
		} else if col == n-1 {
			backtrack(row+1, 0)
		} else {
			backtrack(row, col+1)
		}
	}
	backtrack(0, 0)
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Print(board)
}
