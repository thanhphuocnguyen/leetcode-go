package main

import "fmt"

func minOperations(nums []int) int {
	// increase monotonic stack
	monoSt := make([]int, 0)
	ans := 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		for len(monoSt) > 0 && num < monoSt[len(monoSt)-1] {
			monoSt = monoSt[0 : len(monoSt)-1]
		}
		if len(monoSt) == 0 || num > monoSt[len(monoSt)-1] {
			ans++
			monoSt = append(monoSt, num)
		}
	}

	return ans
}

func main() {
	fmt.Println(minOperations([]int{3, 1, 2, 1}))
	fmt.Println(minOperations([]int{1, 2, 1, 2, 1, 2}))
}
