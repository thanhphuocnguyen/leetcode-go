package main

import "fmt"

func validStrings(n int) []string {
	var backtrack func(s string)
	ans := make([]string, 0)
	backtrack = func(s string) {
		if len(s) == n {
			ans = append(ans, s)
			return
		}
		if len(s) > 0 && s[len(s)-1] == '0' {
			backtrack(s + "1")
		} else {
			backtrack(s + "0")
			backtrack(s + "1")
		}
	}
	backtrack("")
	return ans
}

func main() {
	fmt.Println(validStrings(1))
	fmt.Println(validStrings(3))
}
