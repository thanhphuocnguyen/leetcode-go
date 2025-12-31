package main

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
	var ans int64
	n := len(prices)
	prefixSumBase := make([]int, n+1)
	prefixSumSell := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSumBase[i] = prefixSumBase[i-1] + (prices[i-1] * strategy[i-1])
		prefixSumSell[i] = prefixSumSell[i-1] + prices[i-1]
	}
	ans = int64(prefixSumBase[n])
	for i := k - 1; i < len(prices); i++ {
		base := int64(prefixSumBase[i-k+1]) + int64(prefixSumBase[n]-prefixSumBase[i+1])
		delta := int64(prefixSumSell[i+1] - prefixSumSell[i-k/2+1])
		ans = max(ans, base+delta)
	}

	return ans
}

func main() {
	fmt.Println(maxProfit([]int{4, 2, 8}, []int{-1, 0, 1}, 2))
}
