package main

import (
	"fmt"
	"strconv"
)

func minMaxDifference(num int) int {
	numStr := strconv.Itoa(num)
	var toSwap rune
	for _, b := range numStr {
		if b != '9' {
			toSwap = b
			break
		}
	}
	max, min := 0, 0
	for _, b := range numStr {
		max *= 10
		min *= 10
		if b == toSwap {
			max += 9
		} else {
			max += int(b - '0')
		}
		if b != rune(numStr[0]) {
			min += int(b - '0')
		}
	}
	return max - min
}

func main() {
	fmt.Println(minMaxDifference(90)) // Output: 888
}
