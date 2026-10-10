package main

import "fmt"

func countCommas(n int64) int64 {
	if n < 1000 {
		return 0
	}
	var ans int64 = 0
	if n <= 999_999 {
		return (n - 999)
	}
	ans += 999_999 - 999
	if n > 999_999 && n <= 999_999_999 {
		return ((n - 999_999) * 2) + ans
	}

	ans += 2 * (999_999_999 - 999_999)

	if n > 999_999_999 && n <= 999_999_999_999 {
		return ((n - 999_999_999) * 3) + ans
	}

	ans += 3 * (999_999_999_999 - 999_999_999)

	if n > 999_999_999_999 && n < 999_999_999_999_999 {
		return ((n - 999_999_999_999) * 4) + ans
	}

	ans += 4 * (999_999_999_999_999 - 999_999_999_999)

	if n > 999_999_999_999_999 {
		return ans + 5
	}
	return ans
}

func main() {
	fmt.Println(countCommas(999999999999999))
	fmt.Println(countCommas(365674875645073))
	fmt.Println(countCommas(1000000000000000))
	fmt.Println(countCommas(6225255345))
	fmt.Println(countCommas(1004590))
}
