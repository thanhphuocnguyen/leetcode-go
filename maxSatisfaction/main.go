package main

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	// satisfied & time
	sort.Ints(satisfaction)
	n := len(satisfaction)
	memo := make([][]int, n)
	for i := range n {
		memo[i] = make([]int, n+1)
		for j := range n + 1 {
			memo[i][j] = -1
		}
	}
	return dp(memo, satisfaction, 0, 1)
}

func dp(memo [][]int, s []int, i, time int) int {
	if i == len(memo) {
		return 0
	}
	if memo[i][time] != -1 {
		return memo[i][time]
	}
	rs := 0
	if s[i] >= 0 {
		rs += s[i]*time + dp(memo, s, i+1, time+1)
	} else {
		skip := dp(memo, s, i+1, time)
		take := s[i]*time + dp(memo, s, i+1, time+1)
		rs += max(skip, take)
	}
	memo[i][time] = rs
	return memo[i][time]
}

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -7}))
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
}
