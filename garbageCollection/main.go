package main

import "fmt"

func garbageCollection(garbage []string, travel []int) int {
	n := len(travel)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + travel[i-1]
	}
	ans := 0
	lastSeen := make(map[rune]int)
	for i, gb := range garbage {
		for _, typ := range gb {
			lastSeen[typ] = i
		}
		ans += len(gb)
	}
	for _, gType := range []rune{'G', 'P', 'M'} {
		if _, ok := lastSeen[gType]; ok {
			ans += prefixSum[lastSeen[gType]]
		}
	}
	return ans
}

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
}
