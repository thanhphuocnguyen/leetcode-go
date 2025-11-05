package main

import "fmt"

func rotateTheBox(box [][]byte) [][]byte {
	rows, cols := len(box), len(box[0])
	rotatedBox := make([][]byte, cols)
	for i := range rotatedBox {
		rotatedBox[i] = make([]byte, rows)
	}

	for i := range cols {
		for j := range rows {
			rotatedBox[i][j] = box[rows-j-1][i]
		}
	}

	for col := range rows {
		k := cols - 1
		for row := cols - 1; row >= 0; row-- {
			if rotatedBox[row][col] == '#' {
				for k >= row {
					if rotatedBox[k][col] == '.' {
						rotatedBox[k][col], rotatedBox[row][col] = '#', '.'
						break
					}
					k--
				}
			}
			if rotatedBox[row][col] == '*' {
				k = row - 1
			}
		}
	}
	return rotatedBox
}
func main() {
	box := [][]byte{
		{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'},
	}

	fmt.Println(rotateTheBox(box))
}
