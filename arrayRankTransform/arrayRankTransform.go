package main

import "slices"

func main() {
	arr := []int{40, 10, 20, 30}
	arrayRankTransform(arr)
}

func arrayRankTransform(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	slices.Sort(sortedArr)
	rank := 1
	ranks := make(map[int]int)

	for i, v := range sortedArr {
		if i > 0 && v != sortedArr[i-1] {
			rank++
		}
		ranks[v] = rank
	}

	for i, v := range arr {
		arr[i] = ranks[v]
	}

	return arr
}
