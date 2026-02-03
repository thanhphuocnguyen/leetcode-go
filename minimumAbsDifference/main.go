package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	pairs := make([][]int, 0)
	currMin := math.MaxInt32
	for i, j := 0, 1; j < len(arr); i, j = i+1, j+1 {
		diff := arr[j] - arr[i]
		if diff < currMin {
			currMin = diff
			pairs = pairs[:0]
			pairs = append(pairs, []int{arr[i], arr[j]})
		} else if diff == currMin {
			pairs = append(pairs, []int{arr[i], arr[j]})
		}

	}
	return pairs
}

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
