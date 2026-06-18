package main

import "fmt"

func isCovered(ranges [][]int, left int, right int) bool {
	mini := left
	maxi := right
	for _, r := range ranges {
		mini = min(mini, r[0])
		maxi = max(maxi, r[1])
	}

	n := maxi - mini + 1
	diff := make([]int, n)
	for _, rng := range ranges {
		l, r := rng[0]-mini, rng[1]-mini
		diff[l]++
		if r+1 < n {
			diff[r+1]--
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	for i := left - mini; i < min(right-mini+1, n); i++ {

		if diff[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCovered([][]int{{3, 3}, {1, 1}}, 3, 3))
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
	fmt.Println(isCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
}
