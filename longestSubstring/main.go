package main

import "fmt"

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	freq := [26]int{}
	for _, r := range s {
		freq[r-'a']++
	}
	valid := true
	for _, f := range freq {
		if f != 0 && f < k {
			valid = false
			break
		}
	}
	if valid {
		return len(s)
	}
	i, j := 0, 0
	ans := 0
	for j < len(s) {
		for i < len(s) && freq[s[i]-'a'] < k {
			i++
		}
		for j < len(s) && freq[s[j]-'a'] >= k {
			j++
		}
		if i < j {
			ans = max(ans, longestSubstring(s[i:j], k))
		}
		i = j
		j++
	}
	return ans

}

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("ababbc", 2))
}
