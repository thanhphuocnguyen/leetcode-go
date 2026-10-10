package main

import "fmt"

func maxSubarrayLength(nums []int, k int) int {
	mp := make(map[int]int)
	ans := 0
	i, j := 0, 0
	for j < len(nums) {
		mp[nums[j]]++
		for mp[nums[j]] > k {
			mp[nums[i]]--
			i++
		}
		ans = max(ans, j-i+1)
		j++
	}
	return ans
}

func main() {
	fmt.Println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2))
	fmt.Println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1))
}
