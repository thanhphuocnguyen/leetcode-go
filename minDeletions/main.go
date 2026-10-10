package main

import "fmt"

func minDeletions(s string) int {
	freq := [26]int{}
	for _, c := range s {
		freq[c-'a']++
	}
	freqMp := make(map[int]int)
	// sort.Ints(freq)
	for i := range 26 {
		if freq[i] == 0 {
			continue
		}
		freqMp[freq[i]]++
	}
	ans := 0
	for i := range 26 {
		if freq[i] == 0 {
			continue
		}
		for freqMp[freq[i]] > 1 {
			freqMp[freq[i]]--
			freq[i]--
			freqMp[freq[i]]++
			ans++
			if freq[i] == 0 {
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minDeletions("bbcebab"))
}
