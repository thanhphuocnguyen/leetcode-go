package main

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	prefixSumMat := make([][]int, m)
	for i := range m {
		prefixSumMat[i] = make([]int, n)
	}
	for i, row := range mat {
		for j, num := range row {
			prefixSumMat[i][j] = num
			if i > 0 {
				prefixSumMat[i][j] += prefixSumMat[i-1][j]
			}
			if j > 0 {
				prefixSumMat[i][j] += prefixSumMat[i][j-1]
			}
			if i > 0 && j > 0 {
				prefixSumMat[i][j] -= prefixSumMat[i-1][j-1]
			}
		}
	}
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
	}
	for i := range m {
		for j := range n {
			r1, c1 := max(0, i-k), max(0, j-k)
			r2, c2 := min(i+k, m-1), min(j+k, n-1)
			ans[i][j] = prefixSumMat[r2][c2]
			if r1 > 0 {
				ans[i][j] -= prefixSumMat[r1-1][c2]
			}
			if c1 > 0 {
				ans[i][j] -= prefixSumMat[r2][c1-1]
			}
			if r1 > 0 && c1 > 0 {
				ans[i][j] += prefixSumMat[r1-1][c1-1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2))
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
}
