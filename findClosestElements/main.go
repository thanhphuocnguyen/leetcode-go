package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	max, total := 0, 0
	for i := range k {
		total += abs(arr[i] - x)
	}
	max = total
	idx := 0
	for l, r := 1, k; r < len(arr); l, r = l+1, r+1 {
		a, b := abs(arr[l-1]-x), abs(arr[r]-x)
		total += b - a
		if total < max {
			idx = l
			max = total
		}

	}
	return arr[idx : idx+k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(findClosestElements([]int{1, 1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}
