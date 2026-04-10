package main

import "fmt"

func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	onesRow := make([]int, m)
	onesCol := make([]int, n)
	for i, row := range mat {
		for j, cell := range row {
			if cell == 1 {
				onesRow[i]++
				onesCol[j]++
			}
		}
	}
	ans := 0
	for i, row := range mat {
		if onesRow[i] == 1 {
			for j, cell := range row {
				if cell == 1 && onesCol[j] == 1 {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
}
