package main

import "fmt"

func findMiddleIndex(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	rightPrefixSum := 0
	for i := n - 1; i >= 0; i-- {
		rightPrefixSum += nums[i]
		if rightPrefixSum == prefixSum[i] {
			return i
		}
	}
	return -1

}
func main() {
	fmt.Println(findMiddleIndex([]int{2, 5}))
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4}))
}
