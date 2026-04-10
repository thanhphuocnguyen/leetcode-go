package main

import "fmt"

func findRotation(mat [][]int, target [][]int) bool {
	if check(mat, target) {
		return true
	}
	for i := 0; i < 3; i++ {
		rotated := rotate90(mat)
		if check(rotated, target) {
			return true
		}
		mat = rotated
	}
	return false
}

func rotate90(mat [][]int) [][]int {
	n := len(mat)
	rotatedMat := make([][]int, n)
	for i := range n {
		rotatedMat[i] = make([]int, n)
	}

	for i := range n {
		for j := range n {
			rotatedMat[j][n-i-1] = mat[i][j]
		}
	}
	return rotatedMat
}

func check(mat [][]int, target [][]int) bool {
	for i, row := range mat {
		for j, cell := range row {
			if cell != target[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(findRotation([][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}))
}
