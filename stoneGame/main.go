package main

import "fmt"

func stoneGame(piles []int) bool {
	n := len(piles)
	memo := make([][]int, n)
	for i := range n {
		memo[i] = make([]int, n)
		for j := range n {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == j {
			return piles[i]
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		maxi := max(piles[i]-dp(i+1, j), piles[j]-dp(i, j-1))
		memo[i][j] = maxi
		return memo[i][j]
	}
	return dp(0, n-1) > 0
}

func main() {
	fmt.Println(stoneGame([]int{3, 7, 2, 3}))
}
