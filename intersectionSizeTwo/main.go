package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[1] < b[1]
	})

	// there are will be at least ans equals 2
	ans := 2
	// we know that minPrev will be end - 1 and prevEnd will be end
	prevStart, prevEnd := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		// reset new interval
		if start > prevEnd {
			ans += 2
			prevStart = end - 1
			prevEnd = end
		} else if prevStart < start {
			// if prev start < current start there are two case:
			// if two ends is equal the take the end - 1 to ensure the current interval include in next interval
			// if diff just take the bigger end
			if end == prevEnd {
				prevStart = end - 1
			} else {
				prevStart = end
			}
			if prevStart > prevEnd {
				prevStart, prevEnd = prevEnd, prevStart
			}
			ans++
		}
	}
	return ans
}
func main() {
	fmt.Println(intersectionSizeTwo([][]int{{4, 14}, {6, 17}, {7, 14}, {14, 21}, {4, 7}}))
	fmt.Println(intersectionSizeTwo([][]int{{2, 10}, {3, 7}, {3, 15}, {4, 11}, {6, 12}, {6, 16}, {7, 8}, {7, 11}, {7, 15}, {11, 12}}))
	fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}}))
	fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {3, 7}, {8, 9}}))
	fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}))
	fmt.Println(intersectionSizeTwo([][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}))
}
