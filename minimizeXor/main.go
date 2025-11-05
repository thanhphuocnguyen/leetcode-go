package main

import "fmt"

func minimizeXor(num1 int, num2 int) int {
	return countOneBit(num1)
}

func countOneBit(num int) int {
	cnt := 0
	for num > 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}

func main() {
	fmt.Println(minimizeXor(12, 25))
}
