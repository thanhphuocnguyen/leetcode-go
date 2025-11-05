package main

import "fmt"

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSum := make([][]int, len(matrix))
	for i := range len(matrix) {
		prefixSum[i] = make([]int, len(matrix[0]))
	}
	for i, row := range matrix {
		for j, num := range row {
			prefixSum[i][j] = num
			if i > 0 {
				prefixSum[i][j] += prefixSum[i-1][j]
			}
			if j > 0 {
				prefixSum[i][j] += prefixSum[i][j-1]
			}
			if i > 0 && j > 0 {
				prefixSum[i][j] -= prefixSum[i-1][j-1]
			}
		}
	}
	return NumMatrix{prefixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := this.prefixSum[row2][col2]
	if row1 > 0 {
		ans -= this.prefixSum[row1-1][col2]
	}
	if col2 > 0 {
		ans -= this.prefixSum[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		ans += this.prefixSum[row1-1][col1-1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	ctor := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	fmt.Println(ctor.SumRegion(2, 1, 4, 3))
	fmt.Println(ctor.SumRegion(1, 1, 2, 2))
	fmt.Println(ctor.SumRegion(1, 2, 2, 4))
}
