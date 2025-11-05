package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	pairs1 := make(map[byte]byte)
	pairs2 := make(map[byte]byte)
	// freqS := make(map[byte]int)
	// freqt := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cs, ct := s[i], t[i]
		pairWithCs, ok1 := pairs1[cs]
		pairWithCt, ok2 := pairs2[ct]
		if !ok1 {
			pairs1[cs] = ct
		}
		if !ok2 {
			pairs2[ct] = cs
		}
		if ok1 || ok2 {
			if pairWithCs != ct {
				return false
			}
			if pairWithCt != cs {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("abcdefghijklmnopqrstuvwxyzva", "abcdefghijklmnopqrstuvwxyzck"))
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
