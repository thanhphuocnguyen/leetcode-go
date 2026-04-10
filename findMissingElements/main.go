package main

import (
	"fmt"
	"slices"
	"sort"
)

func findMissingElements(nums []int) []int {
	mini, maxi := 100, 1
	sort.Ints(nums)
	for _, num := range nums {
		mini = min(num, mini)
		maxi = max(num, maxi)
	}
	ans := make([]int, 0)
	for i := mini + 1; i <= maxi-1; i++ {
		_, found := slices.BinarySearch(nums, i)
		if !found {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	fmt.Println(findMissingElements([]int{5, 1}))
	fmt.Println(findMissingElements([]int{7, 8, 6, 9}))
	fmt.Println(findMissingElements([]int{1, 4, 2, 5}))
}
