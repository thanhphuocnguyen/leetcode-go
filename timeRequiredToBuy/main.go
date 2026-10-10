package main

import (
	"fmt"
)

type Node struct {
	kth int
	cnt int
}

func timeRequiredToBuy(tickets []int, k int) int {
	times := 0
	// n := len(tickets)
	// for tickets[k] != 0 {
	// 	mini := math.MaxInt32
	// 	currN := 0
	// 	for i := 0; i < n; i++ {
	// 		if tickets[i] > 0 {
	// 			mini = min(mini, tickets[i])
	// 			currN++
	// 		}
	// 	}
	// 	if mini <= 0 {
	// 		break
	// 	}

	// 	for i := range n {
	// 		if tickets[i] > 0 {
	// 			tickets[i] -= mini
	// 		}
	// 	}

	// 	times += mini * currN

	// }
	i, n := 0, len(tickets)
	for tickets[k] > 0 {
		if tickets[i%n] > 0 {
			times++
		}
		tickets[i%n]--
		i++
	}

	return times
}

func main() {
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	fmt.Println(timeRequiredToBuy([]int{84, 49, 5, 24, 70, 77, 87, 8}, 3))
	fmt.Println(timeRequiredToBuy([]int{15, 66, 3, 47, 71, 27, 54, 43, 97, 34, 94, 33, 54, 26, 15, 52, 20, 71, 88, 42, 50, 6, 66, 88, 36, 99, 27, 82, 7, 72}, 18))
	fmt.Println(timeRequiredToBuy([]int{88, 76, 77, 82, 62, 20, 26, 72, 45, 54}, 2))
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}
