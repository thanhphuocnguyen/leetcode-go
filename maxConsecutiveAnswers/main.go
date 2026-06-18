package main

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	ans := 0
	tCnt, fCnt := 0, 0
	l, r := 0, 0
	for r < len(answerKey) {
		if answerKey[r] == 'T' {
			tCnt++
		} else {
			fCnt++
		}
		minCnt := min(tCnt, fCnt)
		for minCnt > k {
			if answerKey[l] == 'T' {
				tCnt--
			} else {
				fCnt--
			}
			minCnt = min(tCnt, fCnt)
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func main() {
	fmt.Println(maxConsecutiveAnswers("FTFFTFTFTTFTTFTTFFTTFFTTTTTFTTTFTFFTTFFFFFTTTTFTTTTTTTTTFTTFFTTFTFFTTTFFFFFTTTFFTTTTFTFTFFTTFTTTTTTF", 32))
	fmt.Println(maxConsecutiveAnswers("TFFT", 1))
	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1))
}
