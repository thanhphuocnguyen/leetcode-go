package main

import "fmt"

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	rs := make([]int, n)
	for i, num := range nums {
		moves := (num % n) + i
		if num > 0 {
			rs[i] = nums[moves%n]
		} else if num < 0 {
			rs[i] = nums[(moves+n)%n]
		} else {
			rs[i] = num
		}
	}

	return rs
}

func main() {
	fmt.Println(constructTransformedArray([]int{-10, -10}))
	fmt.Println(constructTransformedArray([]int{-1, 4, -1}))
	fmt.Println(constructTransformedArray([]int{3, -2, 1, 1}))
}
