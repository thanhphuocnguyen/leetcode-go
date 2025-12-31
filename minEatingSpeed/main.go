package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	minBnn, maxBnn := 1, 1
	for _, p := range piles {
		maxBnn = max(p, maxBnn)
	}

	var possible = func(k int) bool {
		currH := 0
		for _, p := range piles {
			currH += p / k
			if p%k > 0 {
				currH += 1
			}
		}
		return currH <= h
	}

	for minBnn <= maxBnn {
		mid := minBnn + (maxBnn-minBnn)/2
		if possible(mid) {
			maxBnn = mid - 1
		} else {
			minBnn = mid + 1
		}
	}

	return minBnn
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
