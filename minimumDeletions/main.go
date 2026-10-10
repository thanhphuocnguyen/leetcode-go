package main

import (
	"fmt"
	"math"
)

func minimumDeletions(nums []int) int {
	mini, maxi := math.MaxInt32, math.MinInt32
	n := len(nums)
	minIdx, maxIdx := -1, -1
	for i, num := range nums {
		if num < mini {
			mini = num
			minIdx = i
		}

		if num > maxi {
			maxi = num
			maxIdx = i
		}
	}

	// case remove from front
	idx := max(minIdx, maxIdx)
	front := idx + 1
	// case remove from back
	idx = min(minIdx, maxIdx)
	back := n - idx
	// case remove from both sides
	both := (min(minIdx, maxIdx) + 1) + (n - max(minIdx, maxIdx))

	return min(front, back, both)
}

func main() {
	fmt.Println(minimumDeletions([]int{0, -4, 19, 1, 8, -2, -3, 5}))
	fmt.Println(minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
}
