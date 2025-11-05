package main

func removePalindromeSub(s string) int {
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(s string) bool {
	res := true
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			res = false
			break
		}
	}
	return res
}
