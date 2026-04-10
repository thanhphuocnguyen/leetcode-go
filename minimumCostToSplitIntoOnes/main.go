package main

import "fmt"

func minCost(n int) int {
	// cost memo
	memo := make([]int, n+1)
	return dp(memo, n)
}

func dp(memo []int, x int) int {
	if x == 1 {
		return 0
	}
	if x == 2 {
		return 1
	}
	if memo[x] != 0 {
		return memo[x]
	}
	cost := x / 2 * (x - (x / 2))
	memo[x] = cost + dp(memo, x/2) + dp(memo, x-(x/2))
	return memo[x]
}

func main() {
	fmt.Println(minCost(4))
	fmt.Println(minCost(3))
}
