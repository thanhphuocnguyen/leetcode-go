package main

import "fmt"

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}

func countVowelStrings(n int) int {
	ans := 0
	for _, b := range vowels {
		ans += count(n, b)
	}

	return ans
}

func count(n int, lastC byte) int {
	if n == 1 {
		return 1
	}
	cnt := 0
	for _, b := range vowels {
		if b >= lastC {
			cnt += count(n-1, b)
		}
	}

	return cnt
}

func main() {
	fmt.Println(countVowelStrings(3))
}
