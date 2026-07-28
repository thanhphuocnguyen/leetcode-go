package main

import "fmt"

func countCharacters(words []string, chars string) int {
	freq := [26]byte{}
	for i := 0; i < len(chars); i++ {
		freq[chars[i]-'a']++
	}
	wMap := [26]byte{}
	ans := 0
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			wMap[word[i]-'a']++
		}
		check := true
		for i := 0; i < 26; i++ {
			if wMap[i] > freq[i] {
				check = false
				break
			}
		}
		if check {
			ans += len(word)
		}
		wMap = [26]byte{}
	}
	return ans
}

func main() {
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}
