package main

func shortestCompletingWord(licensePlate string, words []string) string {
	freq := make([]rune, 26)
	for _, r := range licensePlate {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			freq[toLower(r)-'a']++
		}
	}
	ans := ""
	for _, word := range words {
		checkFreq := make([]rune, 26)
		for _, c := range word {
			checkFreq[c-'a']++
		}
		check := true
		for i := range 26 {
			if freq[i] > checkFreq[i] {
				check = false
				break
			}
		}
		if check {
			if len(ans) == 0 || len(word) < len(ans) {
				ans = word
			}
		}
	}
	return ans
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return (r + 'a' - 'A')
	}
	return r
}
