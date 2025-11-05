package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	temp := make([]int, 0)
	for i := range m - 1 {
		for j := 0; i+j < m && j < n; j++ {
			temp = append(temp, mat[i+j][j])
		}
		sort.Ints(temp)
		for j := 0; i+j < m && j < n; j++ {
			mat[j+i][j] = temp[j]
		}
		temp = temp[:0]
	}

	for j := 1; j < n-1; j++ {
		for i := 0; i+j < n && i < m; i++ {
			temp = append(temp, mat[i][i+j])
		}
		sort.Ints(temp)
		for i := 0; i+j < n && i < m; i++ {
			mat[i][i+j] = temp[i]
		}
		temp = temp[:0]
	}

	return mat
}

func main() {

}
