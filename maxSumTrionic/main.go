package main

import "math"

func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	var ans int64 = math.MinInt64
	for i := 1; i < n-2; i++ {
		p, q := i, i
		var midSum int64 = int64(nums[p])
		for q+1 < n && nums[q+1] < nums[q] {
			midSum += int64(nums[q+1])
			q++
		}
		if p == q {
			continue
		}
		l, r := p, q
		var leftSum, rightSum int64 = 0, 0
		var maxLeftSum, maxRightSum int64 = math.MinInt64, math.MinInt64
		for l-1 >= 0 && nums[l-1] < nums[l] {
			leftSum += int64(nums[l-1])
			maxLeftSum = max(maxLeftSum, leftSum)
			l--
		}
		if l == p {
			continue
		}
		for r+1 < n && nums[r+1] > nums[r] {
			rightSum += int64(nums[r+1])
			maxRightSum = max(maxRightSum, rightSum)
			r++
		}
		if r == q {
			continue
		}
		i = r - 1
		ans = max(ans, maxLeftSum+midSum+maxRightSum)
	}

	return ans
}
