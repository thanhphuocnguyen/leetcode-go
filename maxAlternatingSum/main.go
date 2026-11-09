package main

import (
	"fmt"
	"sort"
)

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	for i := range n {
		if nums[i] < 0 {
			nums[i] = -nums[i]
		}
	}
	sort.Ints(nums)
	var ans int64
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a, b := nums[j]*nums[j], nums[i]*nums[i]
		ans += int64(a - b)
	}

	if n%2 != 0 {
		ans += int64(nums[n/2] * nums[n/2])
	}

	return ans
}

func main() {
	fmt.Println(maxAlternatingSum([]int{1, -1, 2, -2, 3, -3}))
}
