package main

import "fmt"

func similarPairs(words []string) int {
	// bitmasking
	mp := make(map[int]int)
	ans := 0
	for _, w := range words {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		if val, ok := mp[mask]; ok {
			ans += val
		}
		mp[mask]++
	}
	return ans
}

func main() {
	fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
}
