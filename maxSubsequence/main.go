package main

import (
	"fmt"
	"sort"
)

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	sortedIndexes := make([][]int, n)
	for i, num := range nums {
		sortedIndexes[i] = []int{i, num}
	}
	sort.Slice(sortedIndexes, func(i, j int) bool {
		return sortedIndexes[i][1] > sortedIndexes[j][1]
	})

	sort.Slice(sortedIndexes[:k], func(i, j int) bool {
		return sortedIndexes[i][0] < sortedIndexes[j][0]
	})
	ans := make([]int, k)
	for i := range k {
		ans[i] = sortedIndexes[i][1]
	}
	return ans
}
func main() {
	fmt.Println(maxSubsequence([]int{33, -27, -9, -83, 48}, 2))
	fmt.Println(maxSubsequence([]int{2, 1, 3, 3}, 2))
	fmt.Println(maxSubsequence([]int{50, -75}, 2))
	fmt.Println(maxSubsequence([]int{3, 4, 3, 3}, 2))
	fmt.Println(maxSubsequence([]int{-1, -2, 3, 4}, 3))
}
