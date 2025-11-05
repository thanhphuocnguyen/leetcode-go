package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
}

func spellchecker(wordlist []string, queries []string) []string {
	vowels := []byte{'u', 'i', 'a', 'o', 'e'}
	isVowel := func(b byte) bool {
		for _, v := range vowels {
			if b == v {
				return true
			}
		}
		return false
	}

	originWordMp := make(map[string]bool)
	caseInsensitiveMp := make(map[string]string)
	ignoreVowelMp := make(map[string]string)

	for _, word := range wordlist {
		originWordMp[word] = true
		wordToLower := strings.ToLower(word)
		if _, ok := caseInsensitiveMp[wordToLower]; !ok {
			caseInsensitiveMp[wordToLower] = word
		}
		bs := []byte(wordToLower)
		for i, b := range bs {
			if isVowel(b) {
				bs[i] = '*'
			}
		}
		transformed := string(bs)
		if _, ok := ignoreVowelMp[transformed]; !ok {
			ignoreVowelMp[transformed] = word
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		qToLower := strings.ToLower(q)
		if _, ok := originWordMp[q]; ok {
			ans[i] = q
		} else if word, ok := caseInsensitiveMp[qToLower]; ok {
			ans[i] = word
		} else {
			bs := []byte(qToLower)
			for j, b := range bs {
				if isVowel(b) {
					bs[j] = '*'
				}
			}
			if val, ok := ignoreVowelMp[string(bs)]; ok {
				ans[i] = val
			}
		}
	}

	return ans
}
