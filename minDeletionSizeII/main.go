package main

import "fmt"

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}
	for i := range n {
		for j := 0; j < i; j++ {
			if isSorted(strs, j, i) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxi := 0
	for _, l := range dp {
		maxi = max(l, maxi)
	}
	return n - maxi
}
func isSorted(strs []string, from, to int) bool {
	for _, str := range strs {
		if str[from] > str[to] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minDeletionSize([]string{"babca", "bbazb"}))
}
