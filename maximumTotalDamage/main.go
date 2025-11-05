package main

import "sort"

func maximumTotalDamage(power []int) int64 {
	freq := make(map[int]int)
	for _, p := range power {
		freq[p]++
	}
	keys := make([]int, len(freq))
	i := 0
	for k := range freq {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	n := len(keys)
	dp := make([]int64, n)
	dp[0] = int64(freq[keys[0]]) * int64(keys[0])
	for i := 1; i < n; i++ {
		key := keys[i]
		take := int64(freq[key]) * int64(key)
		prevIdx := bisectRight(keys, i-1, key-3)
		if prevIdx >= 0 {
			take += dp[prevIdx]
		}
		dp[i] = max(dp[i-1], take)
	}
	return dp[i-1]
}

func bisectRight(arr []int, end, target int) int {
	start, idx := 0, -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] <= target {
			idx = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return idx
}

func main() {
	maximumTotalDamage([]int{1, 2, 3, 4, 5})
}
