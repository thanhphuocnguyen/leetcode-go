package main

import "fmt"

func kMirror(k int, n int) int64 {
	var ans int64
	for l := int64(1); n > 0; l *= 10 {
		for i := l; n > 0 && i < l*10; i++ {
			p := createPalindrome(i, true)
			if isPalindrome(p, k) {
				ans += p
				n--
			}
		}
		for i := l; n > 0 && i < l*10; i++ {
			p := createPalindrome(i, false)
			if isPalindrome(p, k) {
				ans += p
				n--
			}
		}
	}
	return ans
}

func createPalindrome(num int64, odd bool) int64 {
	x := num
	if odd {
		x /= 10
	}
	for x > 0 {
		num = num*10 + x%10
		x /= 10
	}
	return num
}

func isPalindrome(num int64, base int) bool {
	chars := make([]rune, 0)
	if num > 0 {
		chars = append(chars, rune(num%int64(base)))
		num /= int64(base)
	}
	s := string(chars)
	i, j := 0, len(s)-1
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
	fmt.Println(kMirror(5, 2)) // Example usage
}
