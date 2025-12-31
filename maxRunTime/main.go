package main

import (
	"fmt"
	"sort"
)

func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	var total int64
	for _, bat := range batteries {
		total += int64(bat)
	}
	for _, bat := range batteries {
		//
		avg := total / int64(n)
		if int64(bat) > avg {
			total -= int64(bat)
			n -= 1
		}
	}
	return total / int64(n)
}

func main() {
	fmt.Println(maxRunTime(3, []int{18, 54, 2, 53, 87, 31, 71, 4, 29, 25}))
}
