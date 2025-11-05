package main

import (
	"strconv"
	"strings"
)

func slidingPuzzle(board [][]int) int {
	target := "123450"
	var sb strings.Builder
	directions := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	for _, row := range board {
		for _, col := range row {
			sb.WriteString(strconv.Itoa(col))
		}
	}
	moves := 0
	queue := make([]string, 0)
	set := make(map[string]bool)
	queue = append(queue, sb.String())
	set[sb.String()] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			currentStr := queue[0]
			if currentStr == target {
				return moves
			}
			queue = queue[1:]
			zeroIndex := strings.IndexByte(currentStr, '0')
			for _, dir := range directions[zeroIndex] {
				nextStr := swap(currentStr, zeroIndex, dir)
				if _, ok := set[nextStr]; ok {
					continue
				}
				set[nextStr] = true
				queue = append(queue, nextStr)
			}
		}
		moves++
	}
	return -1
}

func swap(str string, i, j int) string {
	bytes := []byte(str)
	bytes[i], bytes[j] = bytes[j], bytes[i]
	return string(bytes)
}

func main() {
	board := [][]int{{1, 2, 3}, {4, 0, 5}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {5, 4, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{4, 1, 2}, {5, 0, 3}}
	println(slidingPuzzle(board))
	board = [][]int{{3, 2, 4}, {1, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {4, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {4, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3, 4}, {5, 0, 6, 7}}
	println(slidingPuzzle(board))
	board = [][]int{{0, 1, 2}, {3, 4, 5}}
	println(slidingPuzzle(board))
	board = [][]int{{3, 2, 4}, {1, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {4, 0, 5}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {5, 4, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{4, 1, 2}, {5, 0, 3}}
	println(slidingPuzzle(board))
	board = [][]int{{3, 2, 4}, {1, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {4, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {4, 5, 0}}
	println(slidingPuzzle(board))
	board = [][]int{{1, 2, 3, 4}, {5, 0, 6, 7}}
	println(slidingPuzzle(board))
	board = [][]int{{0, 1, 2}, {3, 4, 5}}
	println(slidingPuzzle(board))
}
