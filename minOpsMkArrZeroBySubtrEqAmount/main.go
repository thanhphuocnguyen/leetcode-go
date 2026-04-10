package main

import (
	"fmt"
	"sort"
)

func minimumOperations(nums []int) int {
	n, ans := len(nums), 0
	sort.Ints(nums)
	if nums[n-1] == 0 {
		return 0
	}
	reduced := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 || nums[i] == reduced {
			continue
		}

		ans++
		nums[i] -= reduced
		reduced += nums[i]
		nums[i] = 0
	}

	return ans
}

func main() {
	fmt.Println(minimumOperations([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
}
