package main

import "fmt"

func numSplits(s string) int {
	n := len(s)
	seen := make([]bool, 26)
	prefixSumLeft := make([]int, n+1)
	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		if !seen[idx] {
			seen[idx] = true
			prefixSumLeft[i+1] = prefixSumLeft[i] + 1
		} else {
			prefixSumLeft[i+1] = prefixSumLeft[i]
		}
	}
	seen = make([]bool, 26)

	prefixSumRight := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		idx := s[i] - 'a'
		if !seen[idx] {
			seen[idx] = true
			prefixSumRight[i] = prefixSumRight[i+1] + 1
		} else {
			prefixSumRight[i] = prefixSumRight[i+1]
		}
	}
	ans := 0
	for i := 1; i < n; i++ {
		if prefixSumLeft[i] == prefixSumRight[i] {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(numSplits("aacaba"))
	fmt.Println(numSplits("abcd"))
}
