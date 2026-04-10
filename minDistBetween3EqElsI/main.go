package main

import (
	"fmt"
	"math"
)

func minimumDistance(nums []int) int {
	mp := make(map[int][]int)
	ans := math.MaxInt32
	for i, num := range nums {
		if pair, ok := mp[num]; ok {
			if pair[0] != -1 && pair[1] != -1 {
				ans = min(ans, (i-pair[0])+(i-pair[1])+(pair[1]-pair[0]))
			}
			mp[num][0] = mp[num][1]
			mp[num][1] = i
		} else {
			mp[num] = []int{-1, i}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumDistance([]int{1, 1, 2, 3, 2, 1, 2}))
	fmt.Println(minimumDistance([]int{1, 2, 1, 1, 3}))
}
