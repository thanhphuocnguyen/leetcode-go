package main

import "fmt"

func minCost(colors string, neededTime []int) int {
	total := 0
	for _, time := range neededTime {
		total += time
	}

	maxi := neededTime[0]
	totalMaxi := 0
	prev := colors[0]
	for i := 1; i < len(colors); i++ {
		if prev == colors[i] {
			maxi = max(maxi, neededTime[i])
		} else {
			totalMaxi += maxi
			maxi = neededTime[i]
			prev = colors[i]
		}
	}
	totalMaxi += maxi
	return total - totalMaxi
}

func main() {
	fmt.Println(minCost("abc", []int{1, 2, 3}))
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
}
