package main

import "fmt"

func maximumValue(strs []string) int {
	ans := 0
	for _, str := range strs {
		temp := 0
		for i := range str {
			if !isDigit(str[i]) {
				temp = len(str)
				break
			} else {
				temp = temp*10 + int(str[i]-'0')
			}
		}
		ans = max(ans, temp)
	}

	return ans
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	fmt.Println(maximumValue([]string{"alic3", "bob", "3", "4", "00000"}))
	fmt.Println(maximumValue([]string{"1", "01", "001", "0001"}))
}
