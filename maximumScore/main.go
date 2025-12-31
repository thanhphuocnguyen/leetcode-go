package main

import "fmt"

func maximumScore(nums []int) int64 {
	var prefixSum int64 = 0
	for i := 0; i < len(nums)-1; i++ {
		prefixSum += int64(nums[i])
	}
	suffixMin := nums[len(nums)-1]
	var ans int64 = prefixSum - int64(suffixMin)
	for i := len(nums) - 2; i > 0; i-- {
		suffixMin = min(suffixMin, nums[i])
		prefixSum -= int64(nums[i])
		ans = max(ans, int64(prefixSum-int64(suffixMin)))
	}

	return ans
}

func main() {
	fmt.Println(maximumScore([]int{10, -1, 3, -4, -5}))
	fmt.Println(maximumScore([]int{-7, -5, 3}))
	fmt.Println(maximumScore([]int{1, 1}))
}
