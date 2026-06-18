package main

import "fmt"

func numberOfAlternatingGroups(colors []int) int {
	n := len(colors)
	ans := 0
	for i := range n {
		if colors[i] == colors[(i+2)%n] && colors[i] != colors[(i+1)%n] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1}))
}
