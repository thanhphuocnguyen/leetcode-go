package main

import "fmt"

func main() {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	prefixMatrix := make([][]int, m)
	for i := range m {
		prefixMatrix[i] = make([]int, n)
		for j := range n {
			prefixMatrix[i][j] = int(matrix[i][j] - '0')
			if j > 0 && matrix[i][j] == '1' {
				prefixMatrix[i][j] += prefixMatrix[i][j-1]
			}
		}
	}
	ans := 0
	for i := range m {
		for j := range n {
			if prefixMatrix[i][j] == 0 {
				continue
			}
			k := i
			width := prefixMatrix[i][j]
			for k < m && prefixMatrix[k][j] != 0 {
				height := k - i + 1
				width = min(width, prefixMatrix[k][j])
				ans = max(ans, height*width)
				k++
			}
			width = prefixMatrix[i][j]
			k = i
			for k >= 0 && prefixMatrix[k][j] != 0 {
				height := k - i + 1
				width = min(width, prefixMatrix[k][j])
				ans = max(ans, height*width)
				k--
			}
		}
	}

	return ans
}
