package main

import (
	"fmt"
)

func numberOfSpecialChars(word string) int {
	lowerMp := [26]int{}
	upperMp := [26]int{}
	for i := range 26 {
		lowerMp[i] = -1
		upperMp[i] = -1
	}

	ans := 0
	for i, r := range word {
		if isLower(r) {
			lowerMp[r-'a'] = i
		} else if upperMp[r-'A'] == -1 {
			upperMp[r-'A'] = i
		}
	}
	for i, idxLower := range lowerMp {
		if upperMp[i] != -1 && idxLower != -1 {
			if upperMp[i] > idxLower {
				ans++
			}
		}
	}
	return ans
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func main() {
	fmt.Println(numberOfSpecialChars("AbBCab"))
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
	fmt.Println(numberOfSpecialChars("abc"))
}
