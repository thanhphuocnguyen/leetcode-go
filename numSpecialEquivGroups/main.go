package main

import (
	"fmt"
)

func numSpecialEquivGroups(words []string) int {
	ans := 0
	mpOdd := make(map[[52]int]bool)
	for _, word := range words {
		temp := [52]int{}
		for i, r := range word {
			if i%2 == 0 {
				temp[r-'a']++
			} else {
				temp[26+(r-'a')]++
			}
		}
		if !mpOdd[temp] {
			mpOdd[temp] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"aa", "bb", "ab", "ba"}))
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}
