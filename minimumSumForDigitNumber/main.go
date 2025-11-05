package main

import "fmt"

func minimumSum(num int) int {
	var freq [10]int = [10]int{}
	for num > 0 {
		freq[num%10]++
		num /= 10
	}
	ans, i, j := 0, 0, 9
	for i <= j {
		for i <= j && freq[i] == 0 {
			i++
		}
		for j >= i && freq[j] == 0 {
			j--
		}
		if i <= j {
			ans += (i*10 + j)
			freq[j]--
			freq[i]--
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumSum(2932))
}
