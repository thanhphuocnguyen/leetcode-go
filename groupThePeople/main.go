package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	mp := make(map[int][]int)
	ans := make([][]int, 0)
	for id, size := range groupSizes {
		if _, ok := mp[size]; !ok {
			mp[size] = []int{id}
		} else {
			if len(mp[size]) == size {
				ans = append(ans, mp[size])
				mp[size] = []int{}
			}
			mp[size] = append(mp[size], id)
		}
	}
	for _, values := range mp {
		ans = append(ans, values)
	}
	return ans
}

func main() {
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}
