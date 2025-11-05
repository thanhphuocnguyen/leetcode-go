package main

import (
	"fmt"
	"math"
)

func maxProductDifference(nums []int) int {
	max1, max2 := math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max2 = num
		}
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}
	}
	return (max1 * max2) - (min1 * min2)
}

func main() {
	fmt.Println(maxProductDifference([]int{4, 2, 5, 9, 7, 4, 8}))
}
