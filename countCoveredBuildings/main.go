package main

import (
	"fmt"
	"sort"
)

func countCoveredBuildings(n int, buildings [][]int) int {
	mpX := make(map[int][]int)
	mpY := make(map[int][]int)
	for _, b := range buildings {
		x, y := b[0], b[1]
		if mpX[x] == nil {
			mpX[x] = []int{}
		}
		if mpY[y] == nil {
			mpY[y] = []int{}
		}
		mpX[x] = append(mpX[x], y)
		mpY[y] = append(mpY[y], x)
	}
	for _, v := range mpX {
		sort.Ints(v)
	}
	for _, v := range mpY {
		sort.Ints(v)
	}
	ans := 0
	for _, b := range buildings {
		x, y := b[0], b[1]
		if y > mpX[x][0] && y < mpX[x][len(mpX[x])-1] && x > mpY[y][0] && x < mpY[y][len(mpY[y])-1] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countCoveredBuildings(3, [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}}))
}
