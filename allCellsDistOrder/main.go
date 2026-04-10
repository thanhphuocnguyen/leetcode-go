package main

import (
	"fmt"
	"sort"
)

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	ans := make([][]int, rows*cols)
	for i := range rows {
		for j := range cols {
			ans[i*cols+j] = []int{i, j}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return abs(rCenter-ans[i][0])+abs(cCenter-ans[i][1]) < abs(rCenter-ans[j][0])+abs(cCenter-ans[j][1])
	})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
}
