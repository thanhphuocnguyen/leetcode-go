package main

import "fmt"

func secondHighest(s string) int {
	var largest, second rune = -1, -1
	for _, r := range s {
		if !isDigit(r) {
			continue
		}
		d := r - '0'
		if d > largest {
			second = largest
			largest = d
		} else if d > second && d < largest {
			second = d
		}
	}
	return int(second)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func main() {
	fmt.Println(secondHighest("dfa12321afd"))
	fmt.Println(secondHighest("abc1111"))
}
