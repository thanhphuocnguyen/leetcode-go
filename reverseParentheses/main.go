package main

import "fmt"

func reverseParentheses(s string) string {
	st := make([]int, 0)
	bs := []byte(s)
	for i, b := range bs {
		switch b {
		case '(':
			st = append(st, i+1)
		case ')':
			start := st[len(st)-1]
			end := i - 1
			for start < end {
				bs[start], bs[end] = bs[end], bs[start]
				start++
				end--
			}

			st = st[:len(st)-1]
		}
	}
	ans := make([]byte, 0, len(s))
	for _, b := range bs {
		if b == '(' || b == ')' {
			continue
		}
		ans = append(ans, b)
	}
	return string(ans)
}

func main() {
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
}
