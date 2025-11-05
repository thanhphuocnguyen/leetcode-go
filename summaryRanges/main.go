package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	start, end := nums[0], nums[0]
	ans := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr-end == 1 {
			end = curr
		} else if end != start {
			ans = append(ans, fmt.Sprintf("%d->%d", start, end))
			start = curr
			end = curr
		} else {
			ans = append(ans, strconv.Itoa(start))
			start = curr
			end = curr

		}
	}
	if start != end {
		ans = append(ans, fmt.Sprintf("%d->%d", start, end))
	} else {
		ans = append(ans, strconv.Itoa(start))
	}
	return ans
}

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
