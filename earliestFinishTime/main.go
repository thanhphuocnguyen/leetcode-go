package main

import (
	"fmt"
	"math"
)

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	ans := math.MaxInt32
	for i, landTime := range landStartTime {
		pickLand := landTime + landDuration[i]
		for j, waterTime := range waterStartTime {
			pickWater := waterTime + waterDuration[j]
			ans = min(ans, max(pickLand, waterTime)+waterDuration[j])
			ans = min(ans, max(pickWater, landTime)+landDuration[i])
		}
	}
	return ans
}

func main() {
	fmt.Println(earliestFinishTime([]int{99}, []int{59}, []int{99, 54}, []int{85, 20}))
	fmt.Println(earliestFinishTime([]int{5}, []int{3}, []int{1}, []int{10}))
	fmt.Println(earliestFinishTime([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
}
