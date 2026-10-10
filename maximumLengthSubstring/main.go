package main

import "fmt"

func maximumLengthSubstring(s string) int {
	freq := [26]int{}
	ans := 0
	for i, j := 0, 0; j < len(s); j++ {
		freq[s[j]-'a']++
		for freq[s[j]-'a'] > 2 {
			freq[s[i]-'a']--
			i++
		}
		ans = max(ans, j-i+1)
	}
	return ans
}

func main() {
	fmt.Println(maximumLengthSubstring("bcbbbcba"))
	fmt.Println(maximumLengthSubstring("aaaa"))
}
