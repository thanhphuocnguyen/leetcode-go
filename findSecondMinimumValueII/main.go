package main

import (
	"fmt"
	"math"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landFirst := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	waterFirst := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	return min(landFirst, waterFirst)
}

func solve(firstStart []int, firstDur []int, secondStart []int, secondDur []int) int {
	minTime := math.MaxInt32
	for i, time := range firstStart {
		dur := firstDur[i]
		minTime = min(minTime, time+dur)
	}
	ans := math.MaxInt32
	for i, time := range secondStart {
		dur := secondDur[i]
		ans = min(ans, max(time, minTime)+dur)
	}
	return ans
}

func main() {
	fmt.Println(
		earliestFinishTime(
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			[]int{2, 3, 4},
			[]int{1, 2, 3},
		),
	)
}
