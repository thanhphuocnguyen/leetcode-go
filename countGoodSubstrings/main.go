package main

import "fmt"

func countGoodSubstrings(s string) int {
	var freq [26]int
	cnt := 0

	for i, j := 0, 0; j < len(s); j++ {
		right := s[j] - 'a'
		freq[right]++
		for freq[right] > 1 {
			freq[s[i]-'a']--
			i++
		}
		if j-i+1 == 3 {
			cnt++
			freq[s[i]-'a']--
			i++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countGoodSubstrings("aababcabc"))
	fmt.Println(countGoodSubstrings("xyzzaz"))
}
