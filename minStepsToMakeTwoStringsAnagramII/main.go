package main

import "fmt"

func minSteps(s string, t string) int {
	freqS := [26]int{}
	freqT := [26]int{}
	for i := range len(s) {
		freqS[s[i]-'a']++
	}
	for i := range len(t) {
		freqT[t[i]-'a']++
	}

	ans := 0
	for k, v := range freqS {
		// not in t
		if v > 0 && freqT[k] == 0 {
			ans += v
		} else if v > 0 && freqT[k] > 0 {
			ans += abs(v - freqT[k])
		}
	}
	for k, v := range freqT {
		if v > 0 && freqS[k] == 0 {
			ans += v
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minSteps("cotxazilut", "nahrrmcchxwrieqqdwdpneitkxgnt"))
}
