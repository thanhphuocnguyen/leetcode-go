package main

import (
	"fmt"
	"math"
)

const MOD int = 1_000_000_007

func countPermutations(complexity []int) int {
	factorial := func(n int) int {
		ans := 1
		for i := 1; i <= n; i++ {
			ans = (ans * i) % MOD
		}
		return ans
	}

	mini, minCnt := math.MaxInt, 0
	for _, c := range complexity {
		if c < mini {
			mini = c
			minCnt = 1
		} else if c == mini {
			minCnt++
		}
	}
	if mini != complexity[0] || minCnt > 1 {
		return 0
	}
	return factorial(len(complexity) - 1)
}

func main() {
	fmt.Println(countPermutations([]int{155, 437, 368, 168}))
	fmt.Println(countPermutations([]int{3, 3, 3, 4, 4, 4}))
	fmt.Println(countPermutations([]int{1, 2, 3}))
}
