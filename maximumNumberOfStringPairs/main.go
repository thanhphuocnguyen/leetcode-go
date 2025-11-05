package main

import "fmt"

func maximumNumberOfStringPairs(words []string) int {
	// 26 * 26 combination all the lowercase english word with length = 2 can be formed.
	mp := make([]bool, 676)
	ans := 0
	for _, w := range words {
		// if reverse of word contains in map -> increase ans
		if mp[int(w[1]-'a')*26+int(w[0]-'a')] {
			ans++
		}
		// set current word in mp
		mp[int(w[0]-'a')*26+int(w[1]-'a')] = true
	}
	return ans
}

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"aa", "wj", "zp", "df", "xb", "sa", "jw", "pz"}))
}
