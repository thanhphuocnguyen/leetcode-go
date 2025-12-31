package main

import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	memo := make([]int, n+1)
	dp(events, memo, 0)
	return ans
}

func dp(events [][]int, memo []int, currEvent []int, idx int, n int) int {
	if idx >= n {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	// take 
	if currEvent[1] < 
	// skip 
	dp(events, memo, )
}
