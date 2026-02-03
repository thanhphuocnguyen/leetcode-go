package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
}
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		a, b := boxTypes[i], boxTypes[j]
		return a[1] > b[1]
	})
	currSize := truckSize
	ans := 0
	for _, bt := range boxTypes {
		bnum, sz := bt[0], bt[1]
		if bnum > currSize {
			ans += currSize * sz
			break
		}
		ans += bnum * sz
		currSize -= bnum
	}

	return ans
}
