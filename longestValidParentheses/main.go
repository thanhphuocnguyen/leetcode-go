package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	st := make([]int, 0, n)
	st = append(st, -1)
	maxSoFar := 0
	for i, r := range s {
		if r == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) > 0 {
				top := st[len(st)-1]
				maxSoFar = max(maxSoFar, i-top)
			} else {
				// start new segment
				st = append(st, i)
			}
		}
	}
	return maxSoFar
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
}
