package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)
	ans := 0
	cnt := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if cnt == 2 {
			cnt = 0
			continue
		} else {
			cnt++
		}
		ans += cost[i]
	}

	return ans
}

func main() {
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
	fmt.Println(minimumCost([]int{1, 2, 3}))
}
