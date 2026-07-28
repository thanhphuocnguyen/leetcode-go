package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	cnt := 0
	i, j := 0, 1
	for j < n {
		s1, e1 := intervals[i][0], intervals[i][1]
		s2, e2 := intervals[j][0], intervals[j][1]
		if (s2 >= s1 && e2 <= e1) || (s1 >= s2 && e1 <= e2) {
			cnt++
			j++
		} else {
			i = j
			j++
		}
	}
	return n - cnt
}

func main() {
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {1, 5}, {2, 5}, {2, 6}}))
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))
}
