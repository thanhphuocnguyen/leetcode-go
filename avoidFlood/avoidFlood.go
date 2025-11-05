package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)

	rainedOnDays := make(map[int]int)
	dryableLakeDays := make([]int, 0)

	for day, rain := range rains {
		if rain > 0 {
			ans[day] = -1
			if dayRained, ok := rainedOnDays[rain]; ok {
				// binary search to find the day that rained
				dryableDix := sort.SearchInts(dryableLakeDays, dayRained)
				// if no idx to dry lake return empty
				if dryableDix == len(dryableLakeDays) {
					return []int{}
				}
				// set dry late to idx that
				ans[dryableLakeDays[dryableDix]] = rain
				// remove idx out of dry array
				dryableLakeDays = append(dryableLakeDays[:dryableDix], dryableLakeDays[dryableDix+1:]...)
			}
			rainedOnDays[rain] = day

		} else {
			dryableLakeDays = append(dryableLakeDays, day)
		}
	}
	for _, idx := range dryableLakeDays {
		ans[idx] = 1
	}

	return ans
}

func main() {
	fmt.Println(avoidFlood([]int{69, 0, 0, 0, 69}))
	fmt.Println(avoidFlood([]int{1, 2, 0, 1, 2}))
	fmt.Println(avoidFlood([]int{1, 2, 0, 0, 2, 1}))
	fmt.Println(avoidFlood([]int{1, 2, 3, 4}))
}
