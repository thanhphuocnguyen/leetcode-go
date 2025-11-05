package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ans := ""
	cp := columnNumber
	for cp > 0 {
		remainder := cp - 1
		cp--
		c := 'A' + rune(remainder%26)

		ans = string(c) + ans
		cp /= 26
	}
	return ans
}

func main() {
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(2147483647))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(2))
}
