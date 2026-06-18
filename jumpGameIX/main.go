package main

import "fmt"

func maxValue(nums []int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], nums[i])
	}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && prefixMax[i] > suffixMin[i+1] {
			ans[i] = ans[i+1]
		} else {
			ans[i] = prefixMax[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxValue([]int{2, 1, 3}))
	fmt.Println(maxValue([]int{2, 3, 1}))
}
