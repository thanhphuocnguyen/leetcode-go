package main

import "math"

func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	for _, r := range ranks {
		minR = min(r, minR)
	}
	var minTime, maxTime int64 = 1, int64(cars * cars * minR)

	for minTime <= maxTime {
		time := minTime + (maxTime-minTime)/2
		if canRepairInTime(ranks, cars, time) {
			maxTime = time - 1
		} else {
			minTime = time + 1
		}
	}
	return minTime
}

func canRepairInTime(ranks []int, cars int, time int64) bool {
	var cnt int64 = 0
	for _, rank := range ranks {
		cnt += int64(math.Sqrt(float64(time / int64(rank))))
		if cnt >= int64(cars) {
			return true
		}
	}
	return false
}

func main() {
	ranks := []int{4, 2, 3, 1}
	cars := 10
	println(repairCars(ranks, cars))
}
