package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	var backtrack func(subset []int, i int)

	backtrack = func(subset []int, idx int) {
		cp := make([]int, len(subset))
		copy(cp, subset)
		ans = append(ans, cp)
		// if idx == len(nums) {
		// 	return
		// }
		for i := idx; i < len(nums); i++ {
			subset = append(subset, nums[i])
			backtrack(subset, i+1)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack([]int{}, 0)
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
