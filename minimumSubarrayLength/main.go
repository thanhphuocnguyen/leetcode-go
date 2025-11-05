package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	minLen := math.MaxInt32
	start, end, curOrCompound := 0, 0, 0
	for end < len(nums) {
		curOrCompound |= nums[end]
		for start <= end && curOrCompound >= k {
			minLen = min(minLen, end-start+1)
			curOrCompound ^= nums[start]
			start++
		}
		end++
	}
	if minLen == math.MaxInt32 {
		return -1
	}
	return minLen
}

func main() {
	nums := []int{1, 2, 32, 21}
	k := 55
	minimumSubarrayLength(nums, k)
}
