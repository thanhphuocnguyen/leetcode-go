package main

import "fmt"

func toHex(num int) string {
	ans := make([]byte, 0)
	if num < 0 {
		num = 1<<32 + num
	}
	for num > 0 {
		remainder := num % 16
		if remainder > 9 {
			ans = append(ans, byte('a'+(remainder-10)))
		} else {
			ans = append(ans, byte('0'+remainder))
		}
		num /= 16
	}
	for i := range len(ans) / 2 {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return string(ans)
}

func main() {
	fmt.Println(toHex(-1))
	fmt.Println(toHex(26))
}
