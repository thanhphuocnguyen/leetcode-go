package main

import "fmt"

func minStartValue(nums []int) int {
	prefixSum := 0
	mini := nums[0]
	for _, num := range nums {
		prefixSum += num
		mini = min(mini, prefixSum)
	}
	if mini > 0 {
		return 1
	}
	return -mini + 1
}

func main() {
	fmt.Println(minStartValue([]int{2, 3, 5, -5, -1}))
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
}
