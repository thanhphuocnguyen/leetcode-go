package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	ans := ""
	if (numerator < 0) != (denominator < 0) {
		ans += "-"
	}

	var dividend int64 = abs(int64(numerator))
	var divisor int64 = abs(int64(denominator))
	quotient := dividend / divisor

	ans += strconv.Itoa(int(quotient))
	remainder := dividend % divisor
	if remainder == 0 {
		return ans
	}
	ans += "."

	mp := make(map[int64]int)

	for remainder > 0 {
		if _, ok := mp[remainder]; ok {
			pos := mp[remainder]
			temp1, temp2 := ans[:pos], ans[pos:]
			ans = temp1 + "(" + temp2 + ")"
			break
		}
		mp[remainder] = len(ans)
		remainder *= 10
		ans += strconv.Itoa(int(remainder / divisor))
		remainder %= divisor
	}

	return ans
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(4, 9))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(25, 435))
}
