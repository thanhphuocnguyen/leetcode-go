package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := 0
	ans := math.MaxInt32
	for r < len(nums) {
		if sum < target {
			sum += nums[r]
		}
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
