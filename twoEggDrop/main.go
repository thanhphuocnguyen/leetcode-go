package main

import (
	"fmt"
	"math"
)

func twoEggDrop(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for k := 1; k <= i; k++ {
			dp[i] = min(dp[i], 1+max(k-1, dp[i-k]))
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(twoEggDrop(100))
}
