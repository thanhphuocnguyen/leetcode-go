package main

import "fmt"

func numSteps(s string) int {
	n := len(s)
	ans := 0
	var carry byte = '0'
	for i := n - 1; i > 0; i-- {
		b := s[i]
		if carry == '1' {
			if b == '0' {
				carry = '0'
				b = '1'
			} else {
				b = '0'
			}
		}
		if b == '0' {
			ans++
		} else {
			// one of plus 1, one for shift 1
			ans += 2
			carry = '1'
		}
	}
	if carry == '1' {
		return ans + 1
	}
	return ans
}

func main() {
	fmt.Println(numSteps("10"))
	fmt.Println(numSteps("1101"))
	fmt.Println(numSteps("1"))
}

// 1 + 1,
