package main

import "fmt"

func permute(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			cp := make([]int, n)
			copy(cp, nums)
			ans = append(ans, cp)
			return
		}
		for i := idx; i < n; i++ {
			nums[i], nums[idx] = nums[idx], nums[i]
			backtrack(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	backtrack(0)
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
