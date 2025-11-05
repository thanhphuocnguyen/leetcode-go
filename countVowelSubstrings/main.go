package main

import "fmt"

func countVowelSubstrings(word string) int {
	mp := map[byte]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	vowelCnt := 0
	l, r := 0, 0
	ans := 0
	for r < len(word) {
		for !isVowel(word[r]) && l < r {
			if _, ok := mp[word[l]]; ok {
				mp[word[l]]--
				if mp[word[l]] == 0 {
					vowelCnt--
				} else if vowelCnt == 5 {
					ans++
				}
			}
			l++
		}
		if isVowel(word[r]) {
			if mp[word[r]] == 0 {
				vowelCnt++
			}
			if vowelCnt == 5 {
				ans++
			}
			mp[word[r]]++
		}
		r++
	}

	return ans
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	fmt.Println(countVowelSubstrings("cuaieuouac"))
	fmt.Println(countVowelSubstrings("unicornarihan"))
	fmt.Println(countVowelSubstrings("aeiouu"))
}
