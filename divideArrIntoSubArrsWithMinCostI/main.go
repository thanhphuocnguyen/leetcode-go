package main

import "fmt"

func minimumCost(nums []int) int {
	min1, min2 := nums[1], nums[2]
	for i := 2; i < len(nums); i++ {
		n := nums[i]
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
	}
	return nums[0] + min1 + min2
}
func main() {
	fmt.Println(minimumCost([]int{1, 2, 3, 12}))
}
