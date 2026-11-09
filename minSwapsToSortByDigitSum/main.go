package main

import (
	"fmt"
	"sort"
)

func minSwaps(nums []int) int {
	n := len(nums)
	pairs := make([][2]int, n)

	for i, num := range nums {
		pairs[i] = [2]int{num, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		a, b := getSum(pairs[i][0]), getSum(pairs[j][0])
		if a != b {
			return a < b
		}
		return pairs[i][0] < pairs[j][0]
	})

	ans := 0
	target := make([]int, n)
	for i, el := range pairs {
		target[i] = el[1]
	}

	visited := make([]bool, n)

	for i := range n {
		if visited[i] || target[i] == i {
			continue
		}

		j := i
		cycleOfPermutation := 0

		for !visited[j] {
			visited[j] = true
			j = target[j]
			cycleOfPermutation++
		}

		ans += cycleOfPermutation - 1
	}

	return ans
}

func getSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}

	return sum
}

func main() {
	fmt.Println(minSwaps([]int{268835996, 65052660, 415128775}))
}
