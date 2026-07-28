package main

import (
	"fmt"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	n1, n2 := n/2, n/2

	sort.Slice(costs, func(i, j int) bool {
		return abs(costs[i][0]-costs[i][1]) > abs(costs[j][0]-costs[j][1])
	})
	ans := 0
	for _, cost := range costs {
		c1, c2 := cost[0], cost[1]
		if c1 < c2 && n1 > 0 {
			ans += c1
			n1--
		} else if n2 > 0 && c2 < c1 {
			ans += c2
			n2--
		} else if n2 > 0 {
			ans += c2
			n2--
		} else {
			ans += c1
			n1--
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
	fmt.Println(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}))
}
