package main

import "fmt"

func longestBalanced(s string) int {
	n := len(s)
	ans := 0
	for i := range n {
		// freq := make(map[byte]int)
		freq := [26]int{}
		maxFreq := 0
		runeCnt := 0
		eqFreqCnt := 1
		for j := i; j < n; j++ {
			ord := s[j] - 'a'
			freq[ord]++
			if freq[ord] == 1 {
				runeCnt++
			}
			if freq[ord] == maxFreq {
				eqFreqCnt++
			}
			if freq[ord] > maxFreq {
				maxFreq = freq[ord]
				eqFreqCnt = 1
			}
			if eqFreqCnt == runeCnt && j-i+1 > ans {
				ans = j - i + 1
			}
			// freq[s[j]]++
			// mini, maxi := 1000, 1
			// for _, v := range freq {
			// 	mini = min(mini, v)
			// 	maxi = max(maxi, v)
			// }
			// if mini == maxi {
			// 	ans = max(ans, j-i+1)
			// }
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalanced("abbac"))
}
