package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	currMin := points[0]

	ans := 1
	for i := 1; i < len(points); i++ {
		point := points[i]
		if point[0] <= currMin[1] && point[1] >= currMin[0] {
			if point[1] < currMin[1] {
				currMin = point
			}
			continue
		}
		ans += 1
		currMin = point
	}

	return ans
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{0, 9}, {1, 8}, {7, 8}, {1, 6}, {9, 16}, {7, 13}, {7, 10}, {6, 11}, {6, 9}, {9, 13}}))
	fmt.Println(findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}
