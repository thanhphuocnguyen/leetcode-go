package main

import (
	"fmt"
	"sort"
)

func sequentialDigits(low int, high int) []int {
	ans := make([]int, 0)
	for curr := 1; curr < 9; curr++ {
		next := curr + 1

		subTotal := curr
		for (subTotal < low || subTotal <= high) && next < 10 {
			subTotal = subTotal*10 + next
			if subTotal >= low && subTotal <= high {
				ans = append(ans, subTotal)
			}
			next++
		}
	}
	sort.Ints(ans)
	return ans
}

func main() {
	fmt.Println(sequentialDigits(1000, 13000))
	fmt.Println(sequentialDigits(100, 300))
}
