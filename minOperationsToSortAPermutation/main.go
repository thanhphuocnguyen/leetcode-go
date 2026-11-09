package main

import "math"

func minOperations(nums []int) int {
	n := len(nums)
	zeroIdx := -1
	for i, num := range nums {
		if num == 0 {
			zeroIdx = i
			break
		}
	}
	check := func(dir int) bool {
		for i := 1; i < n; i++ {
			prev := (zeroIdx + (i-1)*dir + n) % n
			curr := (zeroIdx + i*dir + n) % n
			if nums[prev] > nums[curr] {
				return false
			}
		}
		return true
	}
	ans := math.MaxInt32
	// case increasing have two case to check,
	// one is rotate left zeroIdx times
	// case reverse and and rotate left n-zeroIdx+1 times
	if check(1) {
		ans = min(ans, zeroIdx)
		ans = min(ans, n-zeroIdx+2)
	}

	// case decreasing also have two cases,
	// one is rotate left zeroIdx+1 times then reverse
	// the other one is reverse then rotate n-zeroIdx-1
	if check(-1) {
		ans = min(ans, zeroIdx+2)
		ans = min(ans, n-zeroIdx)
	}

	return ans
}

func main() {

}
