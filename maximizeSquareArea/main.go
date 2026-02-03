package main

import (
	"fmt"
	"sort"
)

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	const MOD = 1_000_000_007
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	sort.Ints(hFences)
	sort.Ints(vFences)
	mp := make(map[int]int)
	var side int64 = -1
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			mp[hFences[j]-hFences[i]] = -1
		}

	}
	for i := 0; i < len(vFences); i++ {
		for j := i + 1; j < len(vFences); j++ {
			diff := vFences[j] - vFences[i]
			if _, ok := mp[diff]; ok {
				side = max(side, int64(diff))
			}
		}

	}
	if side == -1 {
		return -1
	}

	return int((side * side) % MOD)
}

func main() {
	fmt.Println(maximizeSquareArea(4, 3, []int{2, 3}, []int{2}))
}
