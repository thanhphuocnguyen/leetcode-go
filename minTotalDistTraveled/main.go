package main

import (
	"fmt"
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	m, n := len(robot), len(factory)
	memo := make([][]int64, m)
	for i := range m {
		memo[i] = make([]int64, n)
		for j := range n {
			memo[i][j] = -1
		}
	}
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	return dp(memo, robot, factory, 0, 0)
}

func dp(memo [][]int64, robots []int, factory [][]int, rIdx, fIdx int) int64 {
	// no robot so return 0
	if rIdx == len(robots) {
		return 0
	}

	// no fac so return maxValue
	if fIdx == len(factory) {
		return 1e15
	}

	if memo[rIdx][fIdx] != -1 {
		return memo[rIdx][fIdx]
	}
	// skip
	res := dp(memo, robots, factory, rIdx, fIdx+1)
	var cost int64 = 0
	fac, lim := factory[fIdx][0], factory[fIdx][1]
	for i := 0; i < lim && rIdx+i < len(robots); i++ {
		cost += int64(abs(robots[rIdx+i] - fac))
		res = min(res, cost+dp(memo, robots, factory, rIdx+i+1, fIdx+1))
	}

	memo[rIdx][fIdx] = res
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	fmt.Println(minimumTotalDistance([]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}}))
}
