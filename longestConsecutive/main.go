package main

import "fmt"

func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num] = 1
		prev, prevOk := mp[num-1]
		next, nextOk := mp[num+1]
		if prevOk {
			mp[num] += prev
		} else if nextOk {
			mp[num] += next
		}
	}
	ans := 0
	for _, v := range mp {
		ans = max(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
