package main

import "fmt"

func letterCasePermutation(s string) []string {
	n := len(s)
	var backtrack func(bs []byte, idx int)
	ans := []string{}
	backtrack = func(bs []byte, idx int) {
		ans = append(ans, string(bs))
		for i := idx; i < n; i++ {
			if isDigit(bs[i]) {
				continue
			} else {
				bs[i] ^= 32
				backtrack(bs, i+1)
				bs[i] ^= 32
			}
		}
	}
	backtrack([]byte(s), 0)
	return ans
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	fmt.Println(letterCasePermutation("3z4"))
	fmt.Println(letterCasePermutation("c"))
	fmt.Println(letterCasePermutation("a1b2"))
}
