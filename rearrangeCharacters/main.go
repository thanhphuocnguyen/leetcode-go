package main

import "fmt"

func rearrangeCharacters(s string, target string) int {
	mp1 := make(map[byte]int)
	for i := 0; i < len(target); i++ {
		mp1[target[i]]++
	}
	mp2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := mp1[s[i]]; ok {
			mp2[s[i]]++
		}
	}
	if len(mp2) == 0 {
		return 0
	}
	minCnt := len(s) / len(target)
	for k, v := range mp1 {
		if _, ok := mp2[k]; !ok {
			return 0
		}
		minCnt = min(minCnt, mp2[k]/v)
	}

	return minCnt
}

func main() {
	fmt.Println(rearrangeCharacters("wvu", "tu"))
}
