package main

import "fmt"

func divideString(s string, k int, fill byte) []string {
	n := len(s)
	size := ceilDivision(n, k)
	ans := make([]string, size)
	for i := range size {
		substr := make([]byte, k)
		for j := range k {
			if i*k+j < n {
				substr[j] = s[i*k+j]
			} else {
				substr[j] = fill
			}
		}
		ans[i] = string(substr)
	}
	return ans
}

func ceilDivision(n, k int) int {
	if n%k == 0 {
		return n / k
	}
	return (n / k) + 1
}

func main() {
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}
