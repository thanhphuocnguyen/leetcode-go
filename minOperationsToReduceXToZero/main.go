package main

import "fmt"

func minOperations(nums []int, x int) int {
	n := len(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	subTotal := 0
	currMax := 0
	ans := n + 1
	for i, j := 0, 0; j < n; j++ {
		subTotal += nums[j]
		for subTotal > 0 && total-subTotal < x {
			subTotal -= nums[i]
			i++
		}
		if subTotal >= currMax && (total-subTotal) == x {
			currMax = subTotal
			ans = min(ans, n-(j-i+1))
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{5, 2, 3, 1, 1}, 5))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
}
