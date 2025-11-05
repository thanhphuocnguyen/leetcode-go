package main

import "fmt"

func smallestNumber(pattern string) string {
	initStr := make([]byte, len(pattern)+1)
	for i := 0; i <= len(pattern); i++ {
		initStr[i] = byte('1' + i)
	}
	for !check(string(initStr), pattern) {
		if !findNextPermutation(initStr) {
			break
		}
	}
	return string(initStr)
}

func check(str, pat string) bool {
	for i := 0; i < len(pat); i++ {
		if pat[i] == 'I' && str[i] > str[i+1] {
			return false
		} else if pat[i] == 'D' && str[i] < str[i+1] {
			return false
		}
	}
	return true
}

func findNextPermutation(subsequence []byte) bool {
	lastIncIdx := len(subsequence) - 2
	for lastIncIdx >= 0 && subsequence[lastIncIdx] >= subsequence[lastIncIdx+1] {
		lastIncIdx--
	}
	if lastIncIdx == -1 {
		return false
	}
	swapIdx := len(subsequence) - 1
	for subsequence[swapIdx] <= subsequence[lastIncIdx] {
		swapIdx--
	}
	subsequence[lastIncIdx], subsequence[swapIdx] = subsequence[swapIdx], subsequence[lastIncIdx]
	left, right := lastIncIdx+1, len(subsequence)-1
	for left < right {
		subsequence[left], subsequence[right] = subsequence[right], subsequence[left]
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(smallestNumber("IIIDIDDD"))
}
