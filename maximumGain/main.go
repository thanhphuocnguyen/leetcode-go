package main

import "fmt"

func maximumGain(s string, x int, y int) int {
	higherPoint, lowerPoint := x, y
	higherStr, lowerStr := [2]rune{'a', 'b'}, [2]rune{'b', 'a'}

	if y > x {
		higherPoint = y
		lowerPoint = x
		higherStr, lowerStr = lowerStr, higherStr

	}
	ans := 0
	s, p := removeSubString(s, higherStr, higherPoint)
	ans += p
	_, p = removeSubString(s, lowerStr, lowerPoint)
	ans += p

	return ans
}

func removeSubString(s string, chars [2]rune, point int) (string, int) {
	points := 0
	st := make([]rune, 0)
	for _, c := range s {
		n := len(st)
		if n > 0 && st[n-1] == chars[0] && c == chars[1] {
			st = st[:n-1]
			points += point
		} else {
			st = append(st, c)
		}
	}
	return string(st), points
}

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))
	fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4))
	fmt.Println(maximumGain("aabbrtababbabmaaaeaabeawmvaataabnaabbaaaybbbaabbabbbjpjaabbtabbxaaavsmmnblbbabaeuasvababjbbabbabbasxbbtgbrbbajeabbbfbarbagha", 8484, 4096))
}
