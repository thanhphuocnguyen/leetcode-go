package main

import (
	"fmt"
	"strings"
)

func shortestBeautifulSubstring(s string, k int) string {
	i, j := 0, 0
	freq := make(map[string]bool)
	oneCnt := 0
	mini := len(s)
	for j < len(s) {
		if s[j] == '1' {
			oneCnt++
		}
		for oneCnt >= k {
			if oneCnt == k && j-i+1 <= mini {
				mini = min(mini, j-i+1)
				freq[s[i:j+1]] = true
			}
			if s[i] == '1' {
				oneCnt--
			}
			i++
		}

		if oneCnt == k && j-i+1 <= mini {
			mini = min(mini, j-i+1)
			freq[s[i:j+1]] = true
		}
		j++
	}
	ans := ""
	for k := range freq {
		if len(k) != mini {
			continue
		}
		if ans == "" {
			ans = k
		} else if strings.Compare(k, ans) < 0 {
			ans = k
		}
	}
	return ans
}

func main() {
	fmt.Println(shortestBeautifulSubstring("100011001", 3))
}
