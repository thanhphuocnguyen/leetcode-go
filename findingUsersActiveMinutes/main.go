package main

import "sort"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i][0] != logs[j][0] {
			return logs[i][0] < logs[j][0]
		}
		return logs[i][1] < logs[j][1]
	})
	ans := make([]int, k)
	prvId, cnt := logs[0][0], 0
	for i := 1; i < len(logs); i++ {
		id, time := logs[i][0], logs[i][1]
		if prvId == id && time != logs[i-1][1] {
			cnt++
		} else if prvId != id {
			ans[cnt]++
			cnt = 0
		}
		prvId = id
	}
	ans[cnt]++
	return ans
}

func main() {
	findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5)
}
