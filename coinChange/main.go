package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	memo := make([][]int, len(coins))

	for i := range len(coins) {
		memo[i] = make([]int, amount+1)
		for j := range amount + 1 {
			memo[i][j] = -1
		}
	}
	ans := dp(memo, 0, coins, amount)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func dp(memo [][]int, i int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 || i == len(coins) {
		return math.MaxInt32
	}
	if memo[i][amount] != -1 {
		return memo[i][amount]
	}
	// take
	take := dp(memo, i, coins, amount-coins[i])
	if take != math.MaxInt32 {
		take++
	}
	// skip
	skip := dp(memo, i+1, coins, amount)
	memo[i][amount] = min(take, skip)
	return memo[i][amount]
}

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
