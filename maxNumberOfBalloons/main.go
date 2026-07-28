package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	mp := map[byte]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}
	for i := 0; i < len(text); i++ {
		if _, ok := mp[text[i]]; ok {
			mp[text[i]]++
		}
	}
	minCnt := len(text)
	for k, v := range mp {
		if k == 'l' || k == 'o' {
			minCnt = min(minCnt, v/2)
		} else {
			minCnt = min(minCnt, v)
		}
	}
	return minCnt
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}
