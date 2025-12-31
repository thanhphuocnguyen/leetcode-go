package main

import (
	"fmt"
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	ans := 0
	sort.Ints(apple)
	sort.Ints(capacity)
	i, j := len(apple)-1, len(capacity)-1
	rem := apple[i]
	for i >= 0 && j >= 0 {
		if rem > capacity[j] {
			ans++
			rem -= capacity[j]
			j--
		} else if rem == capacity[j] {
			i--
			j--
			ans++
			if i >= 0 {
				rem = apple[i]
			} else {
				rem = 0
			}
		} else {
			i--
			capacity[j] -= rem
			if i >= 0 {
				rem = apple[i]
			}
		}
	}
	if rem > 0 {
		return ans + 1
	}
	return ans
}

func main() {
	fmt.Println(minimumBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))
	fmt.Println(minimumBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2}))
}
