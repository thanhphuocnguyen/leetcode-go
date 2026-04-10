package main

import "fmt"

func mapWordWeights(words []string, weights []int) string {
	rs := []rune{}
	for _, word := range words {
		sum := 0
		for _, r := range word {
			sum += weights[r-'a']
		}
		rs = append(rs, 'a'+(25-rune(sum%26)))
	}
	return string(rs)
}

func main() {
	fmt.Println(mapWordWeights([]string{"abcd", "def", "xyz"}, []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}))
}
