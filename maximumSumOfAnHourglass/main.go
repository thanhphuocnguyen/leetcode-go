package main

import "fmt"

func maxSum(grid [][]int) int {
	// use prefix sum on 2D array
	// query sum each 3x3 sub matrix
	// remove value at [1,0] and [1,2] cells
	// get max each sub arr and return ans
	m, n := len(grid), len(grid[0])
	prefixSum := make([][]int, m)
	for i := range m {
		prefixSum[i] = make([]int, n)
	}
	for i, row := range grid {
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
	ans := querySum(grid, prefixSum, 0, 0)
	for row1 := 0; row1 < m-2; row1++ {
		for col1 := 0; col1 < n-2; col1++ {
			ans = max(ans, querySum(grid, prefixSum, row1, col1))
		}
	}
	return ans
}

func querySum(grid [][]int, prefixSum [][]int, row1, col1 int) int {
	row2, col2 := row1+2, col1+2
	sum := prefixSum[row2][col2]
	if row1 > 0 {
		sum -= prefixSum[row1-1][col2]
	}
	if col1 > 0 {
		sum -= prefixSum[row2][col1-1]
	}

	if col1 > 0 && row1 > 0 {
		sum += prefixSum[row1-1][col1-1]
	}
	sum -= grid[row2-1][col1] + grid[row2-1][col2]
	return sum
}

func main() {
	fmt.Println(maxSum([][]int{{520626, 685427, 788912, 800638, 717251, 683428}, {23602, 608915, 697585, 957500, 154778, 209236}, {287585, 588801, 818234, 73530, 939116, 252369}}))
	fmt.Println(maxSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(maxSum([][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}}))
}
