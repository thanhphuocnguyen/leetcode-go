package main

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	// merge invervals
	merged := make([][]int, 0)
	for _, meet := range meetings {
		start, end := meet[0], meet[1]
		if len(merged) == 0 || start > merged[len(merged)-1][1] {
			merged = append(merged, meet)
		} else {
			merged[len(merged)-1][1] = max(end, merged[len(merged)-1][1])
		}
	}
	// calculate free days
	ans := merged[0][0] - 1
	for i := 1; i < len(merged); i++ {
		ans += merged[i][0] - merged[i-1][1] - 1
	}
	return ans + (days - merged[len(merged)-1][1])
}

func main() {
	println(countDays(10, [][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}))
	println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
}
