package main

import "sort"

func maximumWeight(intervals [][]int) []int {
	n := len(intervals)
	for i := range n {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	dp := make([][]int64, n+1)
	indices := make([][][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int64, 5)
		indices[i] = make([][]int, 5)
		for j := range 5 {
			indices[i][j] = make([]int, 0)
		}
	}
	for i := range n {
		k := sort.Search(i, func(pos int) bool {
			return intervals[pos][1] >= intervals[i][0]
		})
		for j := 1; j < 5; j++ {
			skip := dp[i][j]
			take := dp[k][j-1] + int64(intervals[i][2])
			if skip > take {
				dp[i+1][j] = dp[i][j]
				indices[i+1][j] = append([]int{}, indices[i][j]...)
				continue
			}

			newIdx := append([]int{}, indices[k][j-1]...)
			newIdx = append(newIdx, intervals[i][3])
			sort.Ints(newIdx)
			if skip == take && compareSlices(indices[i][j], newIdx) < 0 {
				newIdx = append([]int{}, indices[i][j]...)
			}
			dp[i+1][j] = take
			indices[i+1][j] = newIdx
		}
	}
	return indices[n][4]
}

func compareSlices(a, b []int) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := range minLen {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}

	return len(a) - len(b)
}
