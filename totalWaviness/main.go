package main

import "fmt"

func totalWaviness(num1 int, num2 int) int {
	ans := 0
	digits := []int{}
	for i := num1; i <= num2; i++ {
		if i < 100 {
			continue
		}
		x := i
		digits = []int{}
		for x > 0 {
			d := x % 10
			digits = append(digits, d)
			x /= 10
		}
		for i := 1; i+1 < len(digits); i++ {
			if (digits[i] > digits[i-1] && digits[i] > digits[i+1]) || (digits[i] < digits[i-1] && digits[i] < digits[i+1]) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(totalWaviness(4848, 4848))
}
