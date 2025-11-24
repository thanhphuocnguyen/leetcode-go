package main

import (
	"fmt"
)

func shipWithinDays(weights []int, days int) int {
	lo, hi := 0, 0
	for _, num := range weights {
		lo = max(lo, num)
		hi += num
	}

	possible := func(capacity int) bool {
		currDays := 1
		total := 0
		for _, num := range weights {
			if total+num <= capacity {
				total += num
			} else {
				currDays++
				total = num
			}
		}

		return currDays > days
	}

	for lo < hi {
		mid := lo + (hi-lo)/2
		if possible(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
