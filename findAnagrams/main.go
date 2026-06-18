package main

import "fmt"

func findAnagrams(s string, p string) []int {
	pMp := make(map[byte]int)
	for i := range p {
		pMp[p[i]]++
	}
	ans := make([]int, 0)
	l, r := 0, 0
	sMp := make(map[byte]int)
	for r < len(s) {
		if _, ok := pMp[s[r]]; !ok {
			clear(sMp)
			r++
			l = r
			continue
		}
		sMp[s[r]]++
		if r-l+1 == len(p) {
			found := true
			for k, v := range sMp {
				if pMp[k] != v {
					found = false
					break
				}
			}
			if found {
				ans = append(ans, l)
			}
			sMp[s[l]]--
			l++
		}
		r++
	}

	return ans
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
