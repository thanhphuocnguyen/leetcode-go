package main

import "fmt"

func subarraySum(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	ans := 0
	for i := range n {
		ans += prefixSum[i+1] - prefixSum[max(0, i-nums[i])]
	}
	return ans
}

func main() {
	fmt.Println(subarraySum([]int{2, 3, 1}))
}
