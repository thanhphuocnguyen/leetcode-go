package main

import "fmt"

func partition(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)
	isPalindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}

		return true
	}
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(s) && len(path) > 0 {
			cp := make([]string, len(path))
			copy(cp, path)
			ans = append(ans, cp)
			return
		}

		for end := start + 1; end < len(s)+1; end++ {
			ss := s[start:end]
			if isPalindrome(ss) {
				path = append(path, ss)
				backtrack(end)
				if len(path) > 0 {
					path = path[:len(path)-1]
				}
			}
		}
	}
	backtrack(0)
	return ans
}

func main() {
	fmt.Println(partition("aaab"))
	fmt.Println(partition("aab"))
}
