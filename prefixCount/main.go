package main

func prefixCount(words []string, pref string) int {
	lps := constructLPS(pref)
	ans := 0
	for _, word := range words {
		if searchTxt(word, pref, lps) {
			ans++
		}
	}
	return ans
}

func searchTxt(word, pref string, lps []int) bool {
	n, m := len(word), len(pref)
	i, j := 0, 0
	for i < n {
		if word[i] == word[j] {
			i++
			j++
			if j == m {
				return i-j == 0
			}
		} else {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return false
}

func constructLPS(pat string) []int {
	n := len(pat)
	lps := make([]int, n)
	i, curLen := 1, 0
	for i < n {
		if pat[i] == pat[curLen] {
			curLen++
			lps[i] = curLen
			i++
		} else {
			if curLen > 0 {
				curLen = lps[curLen-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func main() {
	words := []string{"apple", "app", "apple", "apples"}
	pref := "app"
	println(prefixCount(words, pref)) // 3
}
