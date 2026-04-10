package main

import "fmt"

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	first, second, third := -1, -1, -1
	for _, row := range grid {
		fmt.Println(row)
	}
	for rowIdx, row := range grid {
		for colIdx, cell := range row {
			if cell > first {
				third = second
				second = first
				first = cell
			} else if cell < first && cell > second {
				third = second
				second = cell
			} else if cell < second && cell > third {
				third = cell
			}

			if rowIdx > 0 && colIdx > 0 && rowIdx < m-1 && colIdx < n-1 {
				for cnt := 1; rowIdx-cnt >= 0 && colIdx-cnt >= 0 && rowIdx+cnt < m && colIdx+cnt < n; cnt++ {
					tempSum := 0
					// decrease row and increase col to calculate top left diagonal
					for i, j := rowIdx, colIdx-cnt; i > rowIdx-cnt; i, j = i-1, j+1 {
						tempSum += grid[i][j]
					}

					// increase row and increase col to calculate top right diagonal
					for i, j := rowIdx-cnt, colIdx; i < rowIdx; i, j = i+1, j+1 {
						tempSum += grid[i][j]
					}

					// increase row and decrease col to calculate bottom right diagonal
					for i, j := rowIdx, colIdx+cnt; i < rowIdx+cnt; i, j = i+1, j-1 {
						tempSum += grid[i][j]
					}

					// decrease row and decrease col to calculate bottom left diagonal
					for i, j := rowIdx+cnt, colIdx; i > rowIdx; i, j = i-1, j-1 {
						tempSum += grid[i][j]
					}
					if tempSum > first {
						third = second
						second = first
						first = tempSum
					} else if tempSum < first && tempSum > second {
						third = second
						second = tempSum
					} else if tempSum < second && tempSum > third {
						third = tempSum
					}
				}
			}

		}
	}

	ans := []int{first}
	if second >= 0 && second != first {
		ans = append(ans, second)
	}
	if third >= 0 && third != second {
		ans = append(ans, third)
	}

	return ans
}

func main() {
	//[[20,17,9,13,5,2,9,1,5],[14,9,9,9,16,18,3,4,12],[18,15,10,20,19,20,15,12,11],[19,16,19,18,8,13,15,14,11],[4,19,5,2,19,17,7,2,2]]
	fmt.Println(getBiggestThree([][]int{{20, 17, 9, 13, 5, 2, 9, 1, 5}, {14, 9, 9, 9, 16, 18, 3, 4, 12}, {18, 15, 10, 20, 19, 20, 15, 12, 11}, {19, 16, 19, 18, 8, 13, 15, 14, 11}, {4, 19, 5, 2, 19, 17, 7, 2, 2}}))
	fmt.Println(getBiggestThree([][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}))
	fmt.Println(getBiggestThree([][]int{{7, 7, 7}}))
	fmt.Println(getBiggestThree([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
