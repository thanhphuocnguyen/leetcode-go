package main

import "fmt"

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	g := make([]int, n)
	f[0] = 1
	g[n-1] = 1
	cnt := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		f[i] = cnt
	}
	cnt = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			cnt++
		} else {
			cnt = 1
		}
		g[i] = cnt
	}
	ans := -1
	for i := 1; i < n; i++ {
		ans = max(ans, min(f[i-1], g[i]))
	}
	return ans
}

func main() {
	fmt.Println(maxIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}))
	fmt.Println(maxIncreasingSubarrays([]int{1, 2, 3, 4, 5}))
}
