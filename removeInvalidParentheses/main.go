package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	rightRemove := 0
	balance := 0
	mp := make(map[string]bool)
	for _, r := range s {
		switch r {
		case '(':
			balance++
		case ')':
			if balance > 0 {
				balance--
			} else {
				rightRemove++
			}
		}
	}
	leftRemove := balance
	ans := make([]string, 0)

	var backtrack func(path []byte, leftCnt, rightCnt, balance, idx int)

	backtrack = func(path []byte, leftCnt, rightCnt, balance, idx int) {
		if idx == len(s) {
			if balance == 0 && leftCnt == 0 && rightCnt == 0 && !mp[string(path)] {
				mp[string(path)] = true
				ans = append(ans, string(path))
			}
			return
		}
		for i := idx; i < len(s); i++ {
			if i > idx && s[i] == s[i-1] {
				continue
			}
			switch s[i] {
			case '(':
				// skip
				if leftCnt > 0 {
					backtrack(path, leftCnt-1, rightCnt, balance, i)
				}
				// take
				path = append(path, s[i])
				backtrack(path, leftCnt, rightCnt, balance+1, i)
				path = path[:len(path)-1]
			case ')':
				// skip
				if rightCnt > 0 {
					backtrack(path, leftCnt, rightCnt-1, balance, i)
				}
				// take
				if balance > 0 {
					path = append(path, s[i])
					backtrack(path, leftCnt, rightCnt, balance-1, i)
					path = path[:len(path)-1]
				}
			default:
				// alphabet characters
				path = append(path, s[i])
				backtrack(path, leftCnt, rightCnt, balance, i)
				path = path[:len(path)-1]
			}
		}

	}

	backtrack([]byte{}, leftRemove, rightRemove, 0, 0)

	return ans
}

func main() {
	fmt.Println(removeInvalidParentheses(")(f"))
	fmt.Println(removeInvalidParentheses("()())()"))
}
