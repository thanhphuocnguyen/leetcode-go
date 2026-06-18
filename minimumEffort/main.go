package main

import (
	"fmt"
	"sort"
)

func minimumEffort(tasks [][]int) int {
	ans := 0
	sort.Slice(tasks, func(i, j int) bool {

		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
	rem := 0
	for _, task := range tasks {
		act, mini := task[0], task[1]
		if rem < mini {
			// add offset between mini and rem
			ans += mini - rem
			rem = mini - act
		} else {
			// we remove energy when task energy accepted
			rem -= act
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumEffort([][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}))
	fmt.Println(minimumEffort([][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}}))
	fmt.Println(minimumEffort([][]int{{1, 2}, {2, 4}, {4, 8}}))
}
