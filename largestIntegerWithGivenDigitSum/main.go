package main

import "fmt"

func largestInteger(n int, s int) int {
	if s == 0 {
		return 0
	}

	if n*9 < s {
		return -1
	}

	ans := 0
	sum := 0
	for i := 0; i < n; i++ {
		for d := 9; d >= 0; d-- {
			if sum+d <= s {
				ans = ans*10 + d
				sum += d
				break
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(largestInteger(2, 17))
}
