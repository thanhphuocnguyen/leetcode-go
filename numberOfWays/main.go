package main

import (
	"fmt"
	"math"
)

const MOD = 10000000007

func numberOfWays(n int, x int) int {
	dp := make([][]int, n+1)
	precomputed := make([]int, n+1)
	for i := 1; i <= n; i++ {
		precomputed[i] = int(math.Pow(float64(i), float64(x)))
	}
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	return backtrack(dp, precomputed, n, x, 0)
}

func backtrack(dp [][]int, precomputed []int, n, x, i int) int {
	if n == 0 {
		dp[n][i] = 1
		return 1
	}
	if dp[n][i] != -1 {
		return dp[n][i]
	}
	sum := 0
	for j := i + 1; precomputed[j] <= n; j++ {
		sum = (sum + backtrack(dp, precomputed, n-precomputed[j], x, j)) % MOD
	}
	dp[n][i] = sum
	return dp[n][i]
}

func main() {
	fmt.Println(numberOfWays(4, 1))
}
