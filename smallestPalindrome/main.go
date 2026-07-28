package main

import "fmt"

func smallestPalindrome(s string) string {
	freq := [26]int{}
	n := len(s)

	for i := range n {
		freq[s[i]-'a']++
	}
	i := 0
	ans := make([]byte, n)

	// fill first half
	for j := 0; j < 26; j++ {

		if freq[j] == 0 {
			continue
		}
		for freq[j] >= 2 {
			b := byte(j + 'a')
			ans[i] = b
			ans[n-i-1] = b
			i++
			freq[j] -= 2
		}
	}
	for i := range 26 {
		if freq[i] > 0 {
			ans[n/2] = byte(i + 'a')
			break
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(smallestPalindrome("babab"))
}
