package main

import "fmt"

func calculateTax(brackets [][]int, income int) float64 {
	prev := 0
	var ans float64 = 0.0
	for _, b := range brackets {
		ic := min(income, b[0]) - prev
		prev = b[0]

		ans += float64(ic) * (float64(b[1]) / 100.0)
	}

	return ans
}

func main() {
	fmt.Println(calculateTax([][]int{{3, 50}, {7, 10}, {12, 25}}, 10))
}
