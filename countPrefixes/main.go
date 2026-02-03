package main

func countPrefixes(words []string, s string) int {
	ans := 0
	for _, w := range words {
		i := 0
		isPrf := true
		for j := range w {
			if w[j] != s[i] {
				isPrf = false
				break
			}
			i++
		}
		if isPrf {
			ans++
		}
	}
	return ans
}
