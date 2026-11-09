package main

import (
	"fmt"
	"sort"
)

func countServers(n int, logs [][]int, x int, queries []int) []int {
	ans := make([]int, len(queries))
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	queriesWithIdx := make([][]int, len(queries))
	for i, q := range queries {
		queriesWithIdx[i] = []int{i, q}
	}
	sort.Slice(queriesWithIdx, func(i, j int) bool {
		return queriesWithIdx[i][1] < queriesWithIdx[j][1]
	})
	active := make([]int, n)
	l, r := 0, 0
	cnt := 0
	for _, q := range queriesWithIdx {
		from, to := q[1]-x, q[1]
		for l < len(logs) && logs[l][1] < from {
			active[logs[l][0]-1]--
			if active[logs[l][0]-1] == 0 {
				cnt--
			}
			l++
		}

		for r < len(logs) && logs[r][1] <= to {
			if active[logs[r][0]-1] == 0 {
				cnt++
			}
			active[logs[r][0]-1]++
			r++
		}
		ans[q[0]] = n - cnt
	}

	return ans
}

func main() {
	fmt.Println(countServers(3, [][]int{{2, 4}, {2, 1}, {1, 2}, {3, 1}}, 2, []int{3, 4}))
	fmt.Println(countServers(3, [][]int{{1, 3}, {2, 6}, {1, 5}}, 5, []int{10, 11}))
}
