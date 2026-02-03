package main

import "fmt"

func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	if words[startIndex] == target {
		return 0
	}
	ans := 101
	for i := startIndex + 1; i < n+startIndex; i++ {
		if words[i%n] == target {
			ans = min(ans, i-startIndex)
		}
	}

	for i := startIndex - 1; i > startIndex-n; i-- {
		if words[(i+n)%n] == target {
			ans = min(ans, startIndex-i)
		}
	}
	if ans == 101 {
		return -1
	}
	return ans
}
func main() {
	fmt.Println(closestTarget([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1))
}
