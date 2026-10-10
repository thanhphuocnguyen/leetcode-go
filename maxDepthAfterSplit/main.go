package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	n := len(seq)
	st := make([]int, 0, n)
	curr := 0
	ans := make([]int, n)
	for i, r := range seq {
		if r == '(' {
			if len(st) > 0 {
				if st[len(st)-1] == 1 {
					curr = 0
				} else {
					curr = 1
				}
			}
			ans[i] = curr
			st = append(st, curr)
		} else {
			if len(st) > 0 {
				ans[i] = st[len(st)-1]
				st = st[:len(st)-1]
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(maxDepthAfterSplit("(()())"))
	fmt.Println(maxDepthAfterSplit("()(())()"))
}
