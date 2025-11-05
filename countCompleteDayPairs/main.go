package main

import "fmt"

func countCompleteDayPairs(hours []int) int {
	mp := make(map[int]int)
	ans := 0
	for _, h := range hours {
		if _, ok := mp[(24-(h%24))%24]; ok {
			ans += mp[(24-(h%24))%24]
		}
		mp[h%24]++
	}
	return ans
}

func main() {
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))
	fmt.Println(countCompleteDayPairs([]int{12, 12, 30, 24, 24}))
}
