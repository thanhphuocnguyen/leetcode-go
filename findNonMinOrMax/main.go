package main

import "fmt"

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	mini, maxi := nums[0], nums[0]
	ans := -1
	for _, num := range nums {
		if num > maxi {
			ans = maxi
			maxi = num
		}
		if num < mini {
			ans = mini
			mini = num
		}
		if num < maxi && num > mini {
			ans = num
		}
	}
	return ans
}

func main() {
	fmt.Println(findNonMinOrMax([]int{3, 30, 24}))
}
