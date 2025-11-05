package main

import "fmt"

func partitionString(s string) int {
	var freq [26]bool
	ans := 1
	for i := range len(s) {
		idx := s[i] - 'a'
		if freq[idx] {
			for i := range 26 {
				freq[i] = false
			}
			freq[idx] = true
			ans++
		}
		freq[idx] = true
	}
	return ans
}

func main() {
	fmt.Println(partitionString("abacaba"))
}
