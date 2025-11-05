package main

import "fmt"

func sortString(s string) string {
	n := len(s)
	var freq [26]int
	for _, r := range s {
		freq[r-'a']++
	}
	bs := make([]byte, n)
	i := 0
	for i < n {
		for j := 0; j < 26; j++ {
			if freq[j] > 0 {
				bs[i] = byte(j + 'a')
				freq[j]--
				i++
			}
		}
		for j := 25; i < n && j >= 0; j-- {
			if freq[j] > 0 {
				bs[i] = byte(j + 'a')
				freq[j]--
				i++
			}
		}
	}
	return string(bs)
}

func main() {
	fmt.Println(sortString("leetcode"))
}
