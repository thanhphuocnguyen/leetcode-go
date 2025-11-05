package main

import "fmt"

// follow up!!
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var boardByteCount = [128]int{}
	for _, row := range board {
		for _, cell := range row {
			boardByteCount[cell]++
		}
	}

	if !checkBoardValid(boardByteCount, word) {
		return false
	}

	if len(word) > 2 && shouldReverseString(boardByteCount, word) {
		word = reverseString(word)
	}
	for i := range m {
		for j := range n {
			if board[i][j] == word[0] && backtrack(board, word, m, n, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func checkBoardValid(boardByteCount [128]int, word string) bool {
	wordByteCnt := [128]int{}
	for i := range len(word) {
		wordByteCnt[word[i]]++
		if wordByteCnt[word[i]] > boardByteCount[word[i]] {
			return false
		}
	}
	return true
}

func shouldReverseString(boardByteCount [128]int, word string) bool {
	firstByte, lastByte := word[0], word[len(word)-1]
	return boardByteCount[lastByte] < boardByteCount[firstByte]
}

func reverseString(word string) string {
	bs := []byte(word)
	for i, j := 0, len(word)-1; i < len(word)/2; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func backtrack(board [][]byte, word string, m, n, i, j, bIdx int) bool {
	if bIdx == len(word) {
		return true
	}

	nextByte := word[bIdx+1]
	temp := board[i][j]
	board[i][j] = '!'
	nextRow, nextCol, prevRow, prevCol := i+1, j+1, i-1, j-1

	if nextRow < m && board[nextRow][j] == nextByte {
		if backtrack(board, word, m, n, nextRow, j, bIdx+1) {
			return true
		}
	}

	if nextCol < n && board[i][nextCol] == nextByte {
		if backtrack(board, word, m, n, i, nextCol, bIdx+1) {
			return true
		}
	}

	if prevRow >= 0 && board[prevRow][j] == nextByte {
		if backtrack(board, word, m, n, prevRow, j, bIdx+1) {
			return true
		}
	}
	if prevCol >= 0 && board[i][prevCol] == nextByte {
		if backtrack(board, word, m, n, i, prevCol, bIdx+1) {
			return true
		}
	}

	board[i][j] = temp
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}
