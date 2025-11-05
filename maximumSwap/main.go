package main

import "strconv"

func main() {
	println(maximumSwap(98368))
}
func maximumSwap(num int) int {
	maxDigitIdx, minDigitIdx, prevMaxDigit := -1, -1, -1
	numstr := strconv.Itoa(num)
	runeArr := make([]rune, len(numstr))
	for i, r := range numstr {
		runeArr[i] = r
	}
	for i := len(numstr) - 1; i >= 0; i-- {
		if maxDigitIdx == -1 || numstr[i] > numstr[maxDigitIdx] {
			maxDigitIdx = i
		} else if numstr[i] < numstr[maxDigitIdx] {
			minDigitIdx = i
			prevMaxDigit = maxDigitIdx
		}
	}

	if prevMaxDigit != -1 && minDigitIdx != -1 {
		runeArr[prevMaxDigit], runeArr[minDigitIdx] = runeArr[minDigitIdx], runeArr[prevMaxDigit]
	}
	ans := 0
	for _, r := range runeArr {
		ans = ans*10 + int(r-'0')
	}

	return ans
}
