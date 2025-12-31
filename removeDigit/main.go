package main

import "fmt"

func removeDigit(number string, digit byte) string {
	rmIdx := -1
	for i := 0; i < len(number); i++ {
		if number[i] != digit {
			continue
		}
		if i+1 < len(number) && number[i] < number[i+1] {
			rmIdx = i
			break
		}
		rmIdx = i
	}
	return number[:rmIdx] + number[rmIdx+1:]
}

func main() {
	fmt.Println(removeDigit("123", '3'))
	fmt.Println(removeDigit("1231", '1'))
	fmt.Println(removeDigit("551", '5'))
}
