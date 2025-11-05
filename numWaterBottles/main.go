package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	total := numBottles
	for total >= numExchange {
		ans += total / numExchange
		canChange := total / numExchange
		remainder := total % numExchange
		total = canChange + remainder
	}
	return ans
}
func main() {
	fmt.Println(numWaterBottles(9, 3))
	fmt.Println(numWaterBottles(15, 4))
}
