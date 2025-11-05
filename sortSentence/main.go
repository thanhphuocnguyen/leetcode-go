package main

import "strings"

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	sorted := make([]string, len(words))
	for _, word := range words {
		idx := word[len(word)-1] - '0' - 1
		sorted[idx] = word[:len(word)-1]
	}
	return strings.Join(sorted, " ")
}

func main() {

}
