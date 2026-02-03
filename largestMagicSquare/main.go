package main

import "fmt"

func main() {
	fmt.Println(largestMagicSquare([][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}))
}

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 1
	isMagicSquare := func(rowIdx, colIdx, k int) bool {
		sum := 0
		for j := colIdx; j < colIdx+k; j++ {
			sum += grid[rowIdx][j]
		}
		for i := rowIdx + 1; i < rowIdx+k; i++ {
			tempSum := 0
			for j := colIdx; j < colIdx+k; j++ {
				tempSum += grid[i][j]
			}
			if sum != tempSum {
				return false
			}
		}

		// cols
		for j := colIdx; j < colIdx+k; j++ {
			tempSum := 0
			for i := rowIdx; i < rowIdx+k; i++ {
				tempSum += grid[i][j]
			}
			if tempSum != sum {
				return false
			}
		}

		// diagonal
		diagonalSum := 0
		for d := 0; d < k; d++ {
			diagonalSum += grid[rowIdx+d][colIdx+d]
		}
		if diagonalSum != sum {
			return false
		}

		diagonalSum = 0
		for d := 0; d < k; d++ {
			diagonalSum += grid[rowIdx+d][colIdx+k-1-d]
		}
		if diagonalSum != sum {
			return false
		}

		return true
	}
	for k := 2; k <= min(m, n); k++ {
		for i := 0; i+k <= m; i++ {
			for j := 0; j+k <= n; j++ {
				if isMagicSquare(i, j, k) {
					ans = k
				}
			}
		}
	}

	return ans
}
