package main

import "fmt"

func countHillValley(nums []int) int {
	prev := nums[0]
	ans := 0
	for i := 1; i < len(nums)-1; i++ {
		num, next := nums[i], nums[i+1]
		if num == next {
			continue
		}
		if (num < prev && num < next) || (num > prev && num > next) {
			ans++
		}
		prev = num
	}
	return ans
}

func main() {
	fmt.Println(countHillValley([]int{8, 2, 5, 7, 7, 2, 10, 3, 6, 2}))
	fmt.Println(countHillValley([]int{6, 5, 10}))
}
