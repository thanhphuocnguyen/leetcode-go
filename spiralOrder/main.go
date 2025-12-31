package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	forward := true
	idx := 0
	row, col := 0, 0
	for row != m && col != n {
		if forward {
			for i := col; i < n; i++ {
				ans[idx] = matrix[row][i]
				idx++
			}
			row++
			for i := row; i < m; i++ {
				ans[idx] = matrix[i][n-1]
				idx++
			}
			n--
		} else {
			for i := n - 1; i >= col; i-- {
				ans[idx] = matrix[m-1][i]
				idx++
			}
			m--
			for i := m - 1; i >= row; i-- {
				ans[idx] = matrix[i][col]
				idx++
			}

			col++
		}
		forward = !forward
	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
