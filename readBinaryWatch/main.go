package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	if turnedOn > 8 || turnedOn < 1 {
		return ans
	}
	var backtrack func(i, cnt, h, m int)
	backtrack = func(i, cnt, h, m int) {
		if cnt == turnedOn {
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			return
		}
		if h > 11 || m > 59 || i == 10 {
			return
		}
		// take
		if i < 4 {
			// backtrack with hour light
			backtrack(i+1, cnt+1, h|(1<<i), m)
		} else {
			// backtrack with minute light
			// minus i with 4 for offset of hours
			backtrack(i+1, cnt+1, h, m|(1<<(i-4)))
		}
		// skip
		backtrack(i+1, cnt, h, m)
	}

	backtrack(0, 0, 0, 0)
	return ans
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
