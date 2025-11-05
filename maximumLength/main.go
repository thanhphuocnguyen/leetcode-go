package main

import "fmt"

type pair struct {
	char      byte
	substrLen int
}

func maximumLength(s string) int {
	cnt := make(map[pair]int)
	for start := range s {
		substrLen := 0
		c := s[start]
		for end := start; end < len(s); end++ {
			if c != s[end] {
				break
			}
			substrLen++
			key := pair{c, substrLen}
			cnt[key]++
		}
	}
	ans := -1
	for key, val := range cnt {
		substrLen := key.substrLen
		if val >= 3 && substrLen > ans {
			ans = substrLen
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength("abcaba")) // 0
}
