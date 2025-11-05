package main

func numberOfSubstrings(s string) int {
	n, left, right := len(s), 0, 0
	abcMap := make(map[byte]int)
	ans := 0
	for right < n {
		abcMap[s[right]]++
		if len(abcMap) == 3 {
			ans += 1 + n - right
		}
		for len(abcMap) == 3 {
			abcMap[s[left]]--
			if abcMap[s[left]] == 0 {
				delete(abcMap, s[left])
			}
			ans += 1 + n - left - right
			left++
		}
		right++
	}

	return ans
}

func main() {
	s := "abcabc"
	println(numberOfSubstrings(s))
}
