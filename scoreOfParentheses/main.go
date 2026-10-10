package main

import "fmt"

func scoreOfParentheses(s string) int {
	st := make([]int, 0, len(s)/2+1)
	// save total all frames
	st = append(st, 0)
	for _, r := range s {
		if r == '(' {
			st = append(st, 0)
		} else {
			currFrame := st[len(st)-1]
			st = st[:len(st)-1]
			outerFrame := st[len(st)-1]
			st = st[:len(st)-1]

			total := outerFrame + max(currFrame*2, 1)
			st = append(st, total)
		}
	}
	return st[0]
}

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
	fmt.Println(scoreOfParentheses("()()"))
	fmt.Println(scoreOfParentheses("((()))"))
}
