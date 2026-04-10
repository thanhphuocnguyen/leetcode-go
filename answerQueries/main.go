package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	ans := make([]int, len(queries))
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	for i, q := range queries {
		idx := binarySearch(prefixSum, q)
		if idx < 0 {
			continue
		}
		ans[i] = idx
	}
	return ans
}

func binarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	fmt.Println(answerQueries(nums, queries))
	fmt.Println(answerQueries([]int{2, 3, 4, 5}, []int{1}))
}
