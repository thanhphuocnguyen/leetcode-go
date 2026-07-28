package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	n := len(nums)
	ans := math.MinInt32
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				ans = max(ans, nums[i]*nums[j]*nums[k])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}))
}
