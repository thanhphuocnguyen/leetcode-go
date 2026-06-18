package main

import (
	"fmt"
	"slices"
)

func solveQueries(nums []int, queries []int) []int {
	numSize := len(nums)
	ans := make([]int, len(queries))
	mp := make(map[int][]int)
	for i, num := range nums {
		if _, ok := mp[num]; !ok {
			mp[num] = []int{i}
		} else {
			mp[num] = append(mp[num], i)
		}
	}

	for i, q := range queries {
		arr := mp[nums[q]]
		n := len(arr)
		if n == 1 {
			ans[i] = -1
			continue
		}

		qIdx, _ := slices.BinarySearch(arr, q)
		// nearest left and right
		l, r := arr[(qIdx-1+n)%n], arr[(qIdx+1)%n]
		ans[i] = min(abs(q-l), numSize-abs(q-l), abs(r-q), numSize-abs(r-q))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(solveQueries([]int{1, 3, 1, 4, 1, 3, 2}, []int{0, 3, 5}))
}
