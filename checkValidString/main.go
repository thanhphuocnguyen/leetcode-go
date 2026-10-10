package main

import "fmt"

func checkValidString(s string) bool {
	// backtrack and dp
	n := len(s)
	memo := make([][]int8, n)
	for i := range n {
		memo[i] = make([]int8, n)
		for j := range n {
			memo[i][j] = -1
		}
	}

	var backtrack func(dp [][]int8, open, i int) bool
	backtrack = func(dp [][]int8, open, i int) bool {
		if open < 0 {
			return false
		}
		if i == len(s) {
			return open == 0
		}

		if dp[i][open] != -1 {
			return dp[i][open] == 1
		}

		ans := false

		if s[i] == '*' {
			ans = ans || backtrack(dp, open+1, i+1) || backtrack(dp, open-1, i+1) || backtrack(dp, open, i+1)
		} else {
			if s[i] == '(' {
				ans = ans || backtrack(dp, open+1, i+1)
			}
			if s[i] == ')' {
				ans = ans || backtrack(dp, open-1, i+1)
			}

		}

		if ans {
			dp[i][open] = 1
		} else {
			dp[i][open] = 0
		}

		return ans
	}
	return backtrack(memo, 0, 0)

	// two stacks
	openParentheses := make([]int, 0, n)
	asterisks := make([]int, 0, n)
	for i, r := range s {
		switch r {
		case '(':
			openParentheses = append(openParentheses, i)
		case ')':
			if len(openParentheses) > 0 {
				openParentheses = openParentheses[:len(openParentheses)-1]
			} else if len(asterisks) > 0 {
				asterisks = asterisks[:len(asterisks)-1]
			} else {
				return false
			}
		default:
			asterisks = append(asterisks, i)
		}
	}

	for len(openParentheses) > 0 && len(asterisks) > 0 {
		if openParentheses[len(openParentheses)-1] > asterisks[len(asterisks)-1] {
			return false
		}
		openParentheses = openParentheses[:len(openParentheses)-1]
		asterisks = asterisks[:len(asterisks)-1]
	}

	return true

	// greedy approach

}

func main() {
	tc1 := "************************************************************"
	fmt.Println(checkValidString(tc1))
	tc1 = "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"
	fmt.Println(checkValidString(tc1))
	tc1 = "(((((*)))**"
	fmt.Println(checkValidString(tc1))
}
