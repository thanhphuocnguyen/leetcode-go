package main

import "fmt"

func minSubarray(nums []int, p int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	target := total % p
	if target == 0 {
		return 0
	}

	// use prefix sum to find where we need to find range of sub array
	// map to track module at index
	mp := make(map[int]int)
	mp[0] = -1
	pref := 0
	cnt := len(nums)
	for i, num := range nums {
		pref = (num + pref) % p
		// need to add p to exclude case pref < target so the rem will be neg
		rem := (pref - target + p) % p
		if _, ok := mp[rem]; ok {
			// get range of prefix sum that sub array where sum eq target
			cnt = min(cnt, i-mp[rem])
		}
		mp[pref] = i
	}
	// if the cnt value eq len of arr so can't remove any sub array
	if cnt == len(nums) {
		return -1
	}
	return cnt
}

func main() {
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))
}
