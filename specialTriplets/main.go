package main

import "fmt"

func specialTriplets(nums []int) int {
	freqPrev := make(map[int]int)
	freqNext := make(map[int]int)
	const MOD int = 1_000_000_007
	for _, num := range nums {
		freqNext[num]++
	}
	ans := 0
	for _, num := range nums {
		freqNext[num]--
		if freqPrev[num*2] > 0 && freqNext[num*2] > 0 {
			ans = (ans + (freqPrev[num*2] * freqNext[num*2])) % MOD
		}
		freqPrev[num]++
	}

	return ans
}

func main() {
	fmt.Println(specialTriplets([]int{8, 4, 2, 8, 4}))
	fmt.Println(specialTriplets([]int{8, 4, 2, 8, 4}))
}
