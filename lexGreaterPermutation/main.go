package main

import "fmt"

func lexGreaterPermutation(s string, target string) string {
	freq := [26]int{}
	for _, c := range s {
		freq[c-'a']++
	}
	ans := make([]byte, 0)
	for i, c := range target {
		pivot := int(c - 'a')
		if freq[pivot] > 0 {
			freq[pivot]--
			if canFormGreater(freq, target, i+1) {
				ans = append(ans, target[i])
				continue
			}
			freq[pivot]++
		}

		for j := pivot + 1; j < 26; j++ {
			if freq[i] > 0 {
				freq[i]--
				ans = append(ans, byte(j+'a'))
				ans = append(ans, getMinStr(freq)...)
				return string(ans)
			}
		}
		return ""
	}
	return ""
}

func canFormGreater(freq [26]int, target string, start int) bool {
	suffix := target[start:]
	largest := getMaxStr(freq)
	return largest > suffix
}

func getMinStr(freq [26]int) []byte {
	bs := make([]byte, 0)
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			for j := 0; j < freq[i]; j++ {
				bs = append(bs, byte(i+'a'))
			}
		}
	}
	return bs
}

func getMaxStr(freq [26]int) string {
	bs := make([]byte, 0)
	for i := 25; i >= 0; i-- {
		if freq[i] > 0 {
			for j := 0; j < freq[i]; j++ {
				bs = append(bs, byte(i+'a'))
			}
		}
	}
	return string(bs)
}
func main() {
	fmt.Println(lexGreaterPermutation("abc", "bba"))
}
