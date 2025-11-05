package main

func longestPalindrome(words []string) int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	currLen, palindromed := 0, 0
	for word, val := range freq {
		if word[0] == word[1] {
			currLen += (val / 2) * 4
			if (val & 1) == 1 {
				palindromed = 1
			}
		} else if word[0] < word[1] {
			reversed := string(word[1]) + string(word[0])
			currLen += min(freq[reversed], val) * 4
		}
	}
	return currLen + palindromed*2
}

func main() {
	// Example usage
	words := []string{"lc", "cl", "gg"}
	result := longestPalindrome(words)
	println(result) // Output: 8
}
