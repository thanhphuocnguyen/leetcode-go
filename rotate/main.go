package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// transpose matrix
	for i := range n {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// reverse row
	for i := 0; i < n; i++ {
		for j := range n / 2 {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(mat)

	fmt.Println(mat)
}
