package main

import (
	"fmt"
	"math"
	"sort"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	subM, subN := m-(k-1), n-(k-1)
	ans := make([][]int, subM)

	for i := range subM {
		ans[i] = make([]int, subN)
	}

	for rowIdx := range subM {
		for colIdx := range subN {
			nums := make([]int, k*k)
			for i := rowIdx; i < rowIdx+k; i++ {
				for j := colIdx; j < colIdx+k; j++ {
					nums[(i-rowIdx)*k+(j-colIdx)] = grid[i][j]
				}
			}
			minAbs := math.MaxInt32
			sort.Ints(nums)
			for i := 1; i < k*k; i++ {
				if nums[i-1] != nums[i] {
					minAbs = min(minAbs, abs(nums[i-1]-nums[i]))
				}
			}

			if minAbs != math.MaxInt32 {
				ans[rowIdx][colIdx] = minAbs
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(minAbsDiff([][]int{{1, -2, 3}, {2, 3, 5}}, 2))
	fmt.Println(minAbsDiff([][]int{{3, -1}}, 1))
	fmt.Println(minAbsDiff([][]int{{1, 8}, {3, -2}}, 2))
}
