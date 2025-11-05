package main

import "sort"

// minimumTotalDistance calculates the minimum total distance for robots to reach factories.
// It takes a slice of integers representing the positions of robots and a slice of slices of integers
// representing the positions and capacities of factories. It returns the minimum total distance as an int64.
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	flattenFactories := []int{}

	for _, fac := range factory {
		for j := 0; j < fac[1]; j++ {
			flattenFactories = append(flattenFactories, fac[0])
		}
	}
	rCnt, ffCnt := len(robot), len(flattenFactories)

	dp := make([][]int64, rCnt+1)
	for i := 0; i < rCnt+1; i++ {
		arr := make([]int64, ffCnt+1)
		dp[i] = arr
	}

	for i := 0; i < rCnt; i++ {
		dp[i][ffCnt] = 1e12
	}

	for i := rCnt - 1; i >= 0; i-- {
		for j := ffCnt - 1; j >= 0; j-- {
			assign := absInt(int64(robot[i]-flattenFactories[j])) + dp[i+1][j+1]
			skip := dp[i][j+1]
			dp[i][j] = min(assign, skip)
		}
	}

	return dp[0][0]
}

func absInt(x int64) int64 {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	robot := []int{0, 4, 6}
	factory := [][]int{{2, 2}, {6, 2}}
	println(minimumTotalDistance(robot, factory))
}
