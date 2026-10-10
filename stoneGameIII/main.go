package main

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n)
	for i := range n {
		dp[i] = math.MinInt32
	}

	var solve func(i int) int
	solve = func(i int) int {
		if i >= n {
			return 0
		}

		if dp[i] != math.MinInt32 {
			return dp[i]
		}

		take := 0
		ans := math.MinInt32
		for k := 0; k < 3 && i+k < n; k++ {
			take += stoneValue[i+k]
			ans = max(ans, take-solve(i+k+1))
		}

		dp[i] = ans
		return ans
	}
	score := solve(0)
	if score > 0 {
		return "Alice"
	} else if score == 0 {
		return "Tie"
	} else {
		return "Bob"
	}
}
