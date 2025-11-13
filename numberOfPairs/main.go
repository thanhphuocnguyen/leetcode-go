package main

import "fmt"

func numberOfPairs(nums []int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	pair, leftOver := 0, 0
	for _, v := range mp {
		if v >= 2 {
			pair += v / 2
		}
		if (v & 1) != 0 {
			leftOver++
		}
	}
	return []int{pair, leftOver}
}

func main() {
	fmt.Println(numberOfPairs([]int{5, 12, 53, 22, 7, 59, 41, 54, 71, 24, 91, 74, 62, 47, 20, 14, 73, 11, 82, 2, 15, 38, 38, 20, 57, 70, 86, 93, 38, 75, 94, 19, 36, 40, 28, 6, 35, 86, 38, 94, 4, 90, 18, 87, 24, 22}))
}
