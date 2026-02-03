package main

import "fmt"

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	ans := 0
	if homePos[0] < startPos[0] {
		for i := startPos[0] - 1; i >= homePos[0]; i-- {
			ans += rowCosts[i]
		}
	}
	if homePos[0] > startPos[0] {
		for i := startPos[0] + 1; i <= homePos[0]; i++ {
			ans += rowCosts[i]
		}
	}

	if homePos[1] < startPos[1] {
		for i := startPos[1] - 1; i >= homePos[1]; i-- {
			ans += colCosts[i]
		}
	}

	if homePos[1] > startPos[1] {
		for i := startPos[1] + 1; i <= homePos[1]; i++ {
			ans += colCosts[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(minCost([]int{1, 0}, []int{2, 3}, []int{5, 4, 3}, []int{8, 2, 6, 7}))
}
