package main

import (
	"fmt"
	"sort"
)

func gcdSum(nums []int) int64 {
	n := len(nums)
	prefixGcd := make([]int, n)
	mxi := 0
	for i, num := range nums {
		mxi = max(mxi, num)
		prefixGcd[i] = gcd(nums[i], mxi)
	}

	sort.Ints(prefixGcd)
	var ans int64

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		ans += int64(gcd(prefixGcd[i], prefixGcd[j]))
	}

	return ans
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	fmt.Println(gcdSum([]int{2, 6, 4}))
}
