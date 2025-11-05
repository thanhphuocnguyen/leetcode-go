package main

import (
	"fmt"
	"sort"
)

var dp [][]int

func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	dp = make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = -1
		}
	}
	return dfs(0, k, events)
}

func dfs(currIndex, count int, events [][]int) int {
	if count == 0 || currIndex == len(events) {
		return 0
	}
	if dp[count][currIndex] != -1 {
		return dp[count][currIndex]
	}
	// skip and pick
	// if pick we add current value and continue with next events that valid in boundary k
	// compare between pick and skip value to fill all dp array and get the largest value at the end of dp
	dp[count][currIndex] = max(dfs(currIndex+1, count, events), events[currIndex][2]+dfs(bisectLeft(events, events[currIndex][1]), count-1, events))
	return dp[count][currIndex]
}

func bisectLeft(events [][]int, target int) int {
	left, right := 0, len(events)
	for left < right {
		mid := (left + right) / 2
		// get event that have start day larger than target (end day of picked event)
		if events[mid][0] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2))

}
