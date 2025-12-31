package main

import "fmt"

func bestClosingTime(customers string) int {
	maxProfit := 0
	profit := 0
	ans := -1
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			profit++
		} else {
			profit--
		}
		if profit > maxProfit {
			maxProfit = profit
			ans = i
		}
	}
	if ans == -1 {
		return 0
	}
	return ans + 1
}
func main() {
	fmt.Println(bestClosingTime("YN"))
	fmt.Println(bestClosingTime("YYYY"))
	fmt.Println(bestClosingTime("NNNNN"))
	fmt.Println(bestClosingTime("YYNY"))
}
