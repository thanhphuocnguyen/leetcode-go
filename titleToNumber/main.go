package main

import "fmt"

func titleToNumber(columnTitle string) int {
	ans := 0
	exponent := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		ans += int(columnTitle[i]-'A'+1) * exponent
		exponent *= 26
	}
	return ans
}

func main() {
	fmt.Println(titleToNumber("XXX"))
}
