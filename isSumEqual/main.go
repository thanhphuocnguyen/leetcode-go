package main

import "fmt"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var source = 0
	sum := 0
	for _, r := range firstWord {
		sum = sum*10 + int(r-'a')

	}
	source += sum
	sum = 0
	for _, r := range secondWord {
		sum = sum*10 + int(r-'a')
	}
	source += sum
	sum = 0
	for _, r := range targetWord {
		sum = sum*10 + int(r-'a')
	}
	return source == sum
}

func main() {
	fmt.Println(isSumEqual("acb", "cba", "cdb"))
}
