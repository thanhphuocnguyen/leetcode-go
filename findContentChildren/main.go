package main

import (
	"fmt"
	"sort"
)

const COUNTING_SORT_THRESHOLD = 100_000

func findContentChildren(g []int, s []int) int {
	maxG, maxS := getMax(g), getMax(s)
	maxNum := max(maxG, maxS)
	if maxNum <= COUNTING_SORT_THRESHOLD {
		countingSort(g, maxG)
		countingSort(s, maxS)
	} else {
		sort.Ints(g)
		sort.Ints(s)
	}

	ans := 0
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			i++
			ans++
		}
	}
	return ans
}

func countingSort(arr []int, maxNum int) {
	count := make([]int, maxNum+1)
	for _, num := range arr {
		count[num]++
	}
	idx := 0
	for i := range maxNum + 1 {
		for count[i] > 0 {
			count[i]--
			arr[idx] = i
			idx++
		}
	}
}

func getMax(arr []int) int {
	ans := 0
	for _, num := range arr {
		ans = max(ans, num)
	}

	return ans
}

func main() {
	// fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{10, 9, 8, 7, 10, 9, 8, 7}, []int{10, 9, 8, 7}))
}
