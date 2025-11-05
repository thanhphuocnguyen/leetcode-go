package main

import "fmt"

func isPalindrome(s string) bool {
	n := len(s)
	start, end := 0, n-1
	for start < end {
		for start < end && !isAlphaNumberic(s[start]) {
			start++
		}
		for start < end && !isAlphaNumberic(s[end]) {
			end--
		}
		if toLower(s[start]) != toLower(s[end]) {
			return false
		}
		start++
		end--
	}
	return true
}
func isAlphaNumberic(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
