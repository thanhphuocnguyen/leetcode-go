package main

import "fmt"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	for i := n - k - 1; i >= 0; i -= 1 {
		energy[i] += energy[i+k]
	}

	ans := energy[0]
	for i := 1; i < n; i++ {
		ans = max(ans, energy[i])
	}
	return ans
}

func main() {
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2))
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3))
}
