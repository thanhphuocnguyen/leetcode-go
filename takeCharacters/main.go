package main

import "fmt"

func takeCharacters(s string, k int) int {
	freq := [3]int{}
	if k == 0 {
		return 0
	}
	for _, r := range s {
		freq[r-'a']++
	}
	for _, v := range freq {
		if v < k {
			return -1
		}
	}
	i, j := 0, 0
	window := [3]int{}
	ans := 0
	for j < len(s) {
		right := s[j] - 'a'
		window[right]++

		for i < j && (freq[0]-window[0] < k || freq[1]-window[1] < k || freq[2]-window[2] < k) {
			left := s[i] - 'a'
			window[left]--
			i++
		}
		if j != i {
			ans = max(ans, j-i+1)
		}
		j++
	}
	return len(s) - ans
}

func main() {
	fmt.Println(takeCharacters("abc", 1))
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
}
