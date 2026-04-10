package main

import "fmt"

func numberOfSubmatrices(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	prefixSumX := make([][]int, m)
	prefixSumY := make([][]int, m)
	ans := 0

	xCntRow := make([]int, m)
	xCntCol := make([]int, n)

	for i, row := range grid {
		prefixSumX[i] = make([]int, n)
		prefixSumY[i] = make([]int, n)
		yCnt := 0
		for j, col := range row {
			switch col {
			case 'X':
				prefixSumX[i][j] = 1
			case 'Y':
				yCnt++
				prefixSumY[i][j] = -1
			}

			if i > 0 {
				xCntRow[i] += xCntRow[i-1]
				prefixSumX[i][j] += prefixSumX[i-1][j]
				prefixSumY[i][j] += prefixSumY[i-1][j]
			}

			if j > 0 {
				xCntCol[j] += xCntCol[j-1]
				prefixSumX[i][j] += prefixSumX[i][j-1]
				prefixSumY[i][j] += prefixSumY[i][j-1]
			}

			if i > 0 && j > 0 {
				prefixSumX[i][j] -= prefixSumX[i-1][j-1]
				prefixSumY[i][j] -= prefixSumY[i-1][j-1]

			}

			if prefixSumX[i][j] > 0 && prefixSumX[i][j]+prefixSumY[i][j] == 0 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numberOfSubmatrices([][]byte{{'.', '.', '.', 'Y'}, {'.', 'Y', '.', 'X'}, {'.', '.', 'X', 'Y'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'.', 'X'}, {'.', 'Y'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'.', '.'}, {'X', 'Y'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'.', '.'}, {'.', '.'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'X'}, {'X', 'Y'}}))
}
