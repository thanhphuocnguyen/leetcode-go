package main

import "fmt"

func removeAnagrams(words []string) []string {
	ans := []string{words[0]}

	var freq [26]int

	for _, curr := range words {
		clearArray(freq)
		prev := ans[len(ans)-1]
		if !isAnagram(prev, curr, freq) {
			ans = append(ans, curr)
		}
	}
	return ans
}

func isAnagram(s1, s2 string, freq [26]int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, r := range s1 {
		freq[r-'a']++
		freq[s2[i]-'a']--
	}

	for i := range 26 {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}

func clearArray(arr [26]int) {
	for i := range len(arr) {
		arr[i] = 0
	}
}

func main() {
	fmt.Println(removeAnagrams([]string{"a", "b", "c", "d", "e"}))
}
