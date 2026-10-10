package main

import "fmt"

var MOD = 1_000_000_007

func numberOfSets(n int, K int) int {
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, K+1)
		dp[i][0] = 1
	}

	for k := 1; k <= K; k++ {
		sum := 0
		for i := 1; i < n; i++ {
			sum = (sum + dp[i-1][k-1]) % MOD
			dp[i][k] = (dp[i-1][k] + sum) % MOD
		}
	}
	return dp[n-1][K]
}

func main() {
	fmt.Println(numberOfSets(4, 2))
}
