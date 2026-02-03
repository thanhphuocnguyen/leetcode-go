package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	i, j := 0, k-1
	ans := math.MaxInt32
	for j < len(nums) {
		ans = min(ans, nums[j]-nums[i])
		j++
		i++
	}
	return ans
}

func main() {
	fmt.Println(minimumDifference([]int{90}, 1))
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
	fmt.Println(minimumDifference([]int{20, 200, 300, 1000}, 3))
}
