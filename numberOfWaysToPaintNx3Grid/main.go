package main

import "fmt"

const MOD int = 1_000_000_007

func numOfWays(n int) int {
	memo := make([][4][4][4]int, n)
	for i := range n {
		memo[i] = [4][4][4]int{}
		for j := range 4 {
			memo[i][j] = [4][4]int{}
			for k := range 4 {
				memo[i][j][k] = [4]int{-1, -1, -1, -1}
			}
		}
	}
	// start with idx 0, and no colors
	return dp(memo, 0, 0, 0, 0)
}

func dp(memo [][4][4][4]int, idx, prev1, prev2, prev3 int) int {
	if idx == len(memo) {
		return 1
	}

	if memo[idx][prev1][prev2][prev3] != -1 {
		return memo[idx][prev1][prev2][prev3]
	}

	ans := 0
	for curr1 := 1; curr1 <= 3; curr1++ {
		if curr1 == prev1 {
			continue
		}
		for curr2 := 1; curr2 <= 3; curr2++ {
			if curr2 == prev2 || curr2 == curr1 {
				continue
			}
			for curr3 := 1; curr3 <= 3; curr3++ {
				if curr3 == prev3 || curr3 == curr2 {
					continue
				}
				ans = (ans + dp(memo, idx+1, curr1, curr2, curr3)) % MOD
			}
		}
	}
	memo[idx][prev1][prev2][prev3] = ans
	return ans
}

func main() {
	fmt.Println(numOfWays(1))
}
