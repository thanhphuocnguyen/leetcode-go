package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	memo := make([][]int, n)
	for i := range n {
		memo[i] = make([]int, 2)
		for j := range 2 {
			memo[i][j] = -1
		}
	}
	return dp(prices, memo, fee, 0, 0)
}

func dp(prices []int, memo [][]int, fee, idx, holding int) int {
	if idx == len(prices) {
		return 0
	}

	if memo[idx][holding] != -1 {
		return memo[idx][holding]
	}

	profit := 0
	if holding == 1 {
		// if holding we will try sell with current stock and try buy new next stock
		sell := prices[idx] - fee + dp(prices, memo, fee, idx, 0)
		// skip current stock and go to try sell at next stock
		skip := dp(prices, memo, fee, idx+1, holding)
		profit = max(profit, sell, skip)
	} else {
		// if do not holding we try buy next stock and flip holding
		buy := -prices[idx] + dp(prices, memo, fee, idx+1, 1)
		// skip buy current stock
		skip := dp(prices, memo, fee, idx+1, holding)
		profit = max(profit, buy, skip)
	}
	memo[idx][holding] = max(memo[idx][holding], profit)
	return memo[idx][holding]
}

func main() {
	fmt.Println(maxProfit([]int{4, 5, 2, 4, 3, 3, 1, 2, 5, 4}, 1))
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}
