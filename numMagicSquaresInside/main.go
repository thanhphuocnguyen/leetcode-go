package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	isMagicSquare := func(rowIdx, colIdx int) bool {
		if grid[rowIdx+1][colIdx+1] != 5 {
			return false
		}
		seen := [10]bool{}
		for i := rowIdx; i < rowIdx+3; i++ {
			for j := colIdx; j < colIdx+3; j++ {
				if grid[i][j] > 9 || grid[i][j] < 0 || seen[grid[i][j]] {
					return false
				}
				seen[grid[i][j]] = true
			}
		}
		sum := grid[rowIdx][colIdx] + grid[rowIdx][colIdx+1] + grid[rowIdx][colIdx+2]

		for i := 1; i < 3; i++ {
			if sum != grid[rowIdx+i][colIdx]+grid[rowIdx+i][colIdx+1]+grid[rowIdx+i][colIdx+2] {
				return false
			}
		}

		for j := 0; j < 3; j++ {
			if sum != grid[rowIdx][colIdx+j]+grid[rowIdx+1][colIdx+j]+grid[rowIdx+2][colIdx+j] {
				return false
			}
		}

		if grid[rowIdx][colIdx]+grid[rowIdx+1][colIdx+1]+grid[rowIdx+2][colIdx+2] != sum {
			return false
		}

		if grid[rowIdx][colIdx+2]+grid[rowIdx+1][colIdx+1]+grid[rowIdx+2][colIdx] != sum {
			return false
		}

		return true
	}
	ans := 0
	for i := range len(grid) - 2 {
		for j := range len(grid[0]) - 2 {
			if isMagicSquare(i, j) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
}
