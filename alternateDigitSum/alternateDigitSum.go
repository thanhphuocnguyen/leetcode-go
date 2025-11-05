package main

func alternateDigitSum(n int) int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	ans := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if (i & 1) != 0 {
			ans -= digits[i]
		} else {
			ans += digits[i]
		}
	}
	return ans
}

func main() {
	alternateDigitSum(10)
}
