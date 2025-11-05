package main

import "fmt"

func minTimeToType(word string) int {
	ans := 0
	var curr rune = 'a'
	for _, r := range word {
		diff := abs(r - curr)
		if diff > 13 {
			ans += int(26 - diff)
		} else {
			ans += int(diff)
		}
		curr = r
		ans++
	}
	return ans
}
func abs(x rune) rune {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(minTimeToType("bza"))
}
