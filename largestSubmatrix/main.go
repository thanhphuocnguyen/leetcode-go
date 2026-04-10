package main

import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	ans := 0
	m, n := len(matrix), len(matrix[0])
	for i := range m {
		for j := range n {
			if i > 0 && matrix[i][j] == 1 {
				// calculate consecutive ones
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	for i := range m {
		// sort asc current row
		sort.Ints(matrix[i])
		for j := range n {
			// min height of current row * width (from j -> end)
			ans = max(ans, matrix[i][j]*(n-j))
		}
	}
	return ans
}

func main() {
	fmt.Println(largestSubmatrix([][]int{{1, 1, 0}, {1, 0, 1}}))
	fmt.Println(largestSubmatrix([][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}))
}
