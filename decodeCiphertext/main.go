package main

import (
	"fmt"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)

	mat := make([][]byte, rows)
	cols := n / rows
	if cols == 0 {
		return ""
	}
	endAtRow, endAtCol := -1, -1
	for i := range rows {
		mat[i] = make([]byte, cols)
		for j := range cols {
			mat[i][j] = encodedText[i*cols+j]
			if encodedText[i*cols+j] != ' ' {
				endAtRow, endAtCol = i, j
			}
		}
	}
	// get original text
	var sb strings.Builder
	for j := range cols {
		for i := range rows {
			if j+i < cols {
				sb.WriteByte(mat[i][j+i])
			}
			if i == endAtRow && i+j == endAtCol {
			}
		}
	}
	return strings.TrimRight(sb.String(), " ")
}

func main() {
	fmt.Println(decodeCiphertext("   a", 2))
	fmt.Println(decodeCiphertext(" b  ac", 2))
	fmt.Println(decodeCiphertext("", 5))
	fmt.Println(decodeCiphertext("iveo    eed   l te   olc", 4))
	fmt.Println(decodeCiphertext("ch   ie   pr", 3))
}
