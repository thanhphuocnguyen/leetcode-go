package main

import "fmt"

func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		square := i * i
		if canPartition(square, i) {
			ans += square
		}
	}
	return ans
}

func canPartition(num, target int) bool {
	if num == 0 || target < 0 {
		return false
	}
	if num == target {
		return true
	}
	return canPartition(num/10, target-num%10) || canPartition(num/100, target-num%100) || canPartition(num/1000, target-num%1000)
}

func main() {
	fmt.Println(punishmentNumber(10))
}
