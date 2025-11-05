package main

import "slices"

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range n {
		dp[i] = 1
		prev[i] = -1
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if groups[i] != groups[j] && dp[i] < dp[j]+1 && check(words[i], words[j]) {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIdx] {
			maxIdx = i
		}
	}
	ans := make([]string, 0)
	for i := maxIdx; i != -1; i = prev[i] {
		ans = append(ans, words[i])
	}
	slices.Reverse(ans)
	return ans
}

func check(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	diff := 0
	for i, c := range w1 {
		if c != rune(w2[i]) {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

func main() {
	println(getWordsInLongestSubsequence([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
}
