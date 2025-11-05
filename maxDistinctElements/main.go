package main

import (
	"fmt"
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	prev := math.MinInt32
	ans := 0
	for _, num := range nums {
		curr := min(num+k, max(prev+1, num-k))
		if curr > prev {
			ans++
			prev = curr
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDistinctElements([]int{56, 56, 54, 54}, 1))
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4}, 1))
	fmt.Println(maxDistinctElements([]int{1, 1, 1, 1, 1, 1, 1, 1, 5, 5, 5}, 3))
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2))
}
