package main

import (
	"fmt"
)

func minSwaps(grid [][]int) int {
	n := len(grid)
	trailingZeros := make([]int, n)
	for i := range grid {
		cnt := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				break
			}
			cnt++
		}
		trailingZeros[i] = cnt
	}
	ans := 0
	for i := range grid {
		if trailingZeros[i] >= n-i-1 {
			continue
		}
		found := false
		j := i + 1
		for ; j < n; j++ {
			if trailingZeros[j] >= n-i-1 {
				found = true
				break
			}
		}

		if !found {
			return -1
		}
		for k := j; k > i; k-- {
			trailingZeros[k], trailingZeros[k-1] = trailingZeros[k-1], trailingZeros[k]
			ans++
		}

	}
	return ans
}

func main() {
	fmt.Println(minSwaps([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}))
	fmt.Println(minSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
}
