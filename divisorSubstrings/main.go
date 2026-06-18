package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	str := strconv.Itoa(num)
	j := 0
	buffer := 1
	for i := 0; i < k-1; i++ {
		buffer *= 10
	}
	curr := 0
	for j < k {
		curr = curr*10 + int(str[j]-'0')
		j++
	}
	ans := 0
	for j < len(str) {
		if curr != 0 && num%curr == 0 {
			ans++
		}
		curr = curr % buffer
		curr = curr*10 + int(str[j]-'0')
		j++
	}
	if curr != 0 && num%curr == 0 {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(divisorSubstrings(430043, 2))
	fmt.Println(divisorSubstrings(240, 2))
}
