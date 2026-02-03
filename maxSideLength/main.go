package main

import "fmt"

func main() {
	fmt.Println(maxSideLength([][]int{{0}}, 1))
	fmt.Println(maxSideLength([][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 1))
	fmt.Println(maxSideLength([][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 4))
}

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	prfSum := make([][]int, m)
	for i, row := range mat {
		prfSum[i] = make([]int, n)
		for j, num := range row {
			prfSum[i][j] = num
			if i > 0 {
				prfSum[i][j] += prfSum[i-1][j]
			}
			if j > 0 {
				prfSum[i][j] += prfSum[i][j-1]
			}
			if i > 0 && j > 0 {
				prfSum[i][j] -= prfSum[i-1][j-1]
			}
		}
	}

	canForm := func(k int) bool {
		// go from first k window -> m
		for row2 := k - 1; row2 < m; row2++ {
			for col2 := k - 1; col2 < n; col2++ {
				row1, col1 := row2-k+1, col2-k+1
				// calculate sum from prefix sum
				sum := prfSum[row2][col2]
				if row1 > 0 {
					sum -= prfSum[row1-1][col2]
				}
				if col1 > 0 {
					sum -= prfSum[row2][col1-1]
				}
				if row1 > 0 && col1 > 0 {
					sum += prfSum[row1-1][col1-1]
				}
				if sum <= threshold {
					return true
				}
			}
		}

		return false
	}

	lo, hi := 1, min(m, n)
	// we use binary search for a window k from 1 -> min(m,n)
	ans := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canForm(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return ans
}
