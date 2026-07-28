package main

import "fmt"

func validPalindrome(s string) bool {
	n := len(s)
	i, j := 0, n-1

	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			skipLeft := isPalindrome(s, i+1, j)
			skipRight := isPalindrome(s, i, j-1)
			if skipRight || skipLeft {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abc"))
	fmt.Println(validPalindrome("eceec"))
	fmt.Println(validPalindrome("lcupuupucul"))
}
