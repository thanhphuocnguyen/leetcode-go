package main

import "fmt"

func longestContinuousSubstring(s string) int {
	ans, cnt := 0, 0
	var start rune = rune(s[0])
	for _, r := range s {
		if r == start {
			cnt++
			start += 1
		} else {
			ans = max(ans, cnt)
			start = r
			cnt = 1
		}
	}
	return max(ans, cnt)
}

func main() {
	fmt.Println(longestContinuousSubstring("yrpjofyzubfsiypfre"))
}
