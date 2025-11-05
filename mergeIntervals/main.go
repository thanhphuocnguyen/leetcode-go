package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	for _, itv := range intervals {
		if len(ans) == 0 || ans[len(ans)-1][1] < itv[0] {
			ans = append(ans, itv)
		} else {
			ans[len(ans)-1][1] = max(itv[1], ans[len(ans)-1][1])
		}
	}
	return ans
}

func main() {
	println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
