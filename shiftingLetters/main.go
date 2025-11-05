package main

import (
	"fmt"
	"strings"
)

func shiftingLetters(s string, shifts [][]int) string {
	prefixShift := make([]int, len(s)+1)
	// init prefixShift
	for _, shift := range shifts {
		l, r, d := shift[0], shift[1], shift[2]
		x := 1
		if d == 0 {
			x = -1
		}
		prefixShift[l] += x
		prefixShift[r+1] -= x
	}
	// compute prefixShift
	numberOfShift := 0
	for i := range len(s) {
		numberOfShift += prefixShift[i]
		prefixShift[i] = numberOfShift
	}
	var sb strings.Builder
	for i, c := range s {
		shiftedChar := 'a' + ((int(c-'a')+prefixShift[i])%26+26)%26
		sb.WriteRune(rune(shiftedChar))
	}
	return sb.String()
}

func main() {
	s := "abc"
	shifts := [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}
	fmt.Println(shiftingLetters(s, shifts))
}
