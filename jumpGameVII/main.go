package main

import "fmt"

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)

	prefixSum := make([]int, n)
	dp := make([]bool, n)
	// can reach '0'
	prefixSum[0] = 1
	dp[0] = true
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			prefixSum[i] = prefixSum[i-1]
			continue
		}
		l := i - maxJump
		r := i - minJump
		sumLeft := 0
		if l > 0 {
			sumLeft = prefixSum[l-1]
		}
		if i >= minJump && s[i] == '0' {
			sumRight := prefixSum[r]
			countReachable := sumRight - sumLeft
			dp[i] = countReachable > 0
		}
		prefixSum[i] = prefixSum[i-1]
		if dp[i] {
			prefixSum[i]++
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(canReach("0000000000", 2, 5))
	fmt.Println(canReach("01", 1, 1))
	fmt.Println(canReach("011010", 2, 3))
}
