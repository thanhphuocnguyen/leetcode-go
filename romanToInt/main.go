package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV")) // 3
}

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sLen := len(s) - 1
	sum := 0
	for i := sLen; i >= 0; i-- {
		if i < sLen && roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}
	return sum
}
