package main

import "fmt"

func dominantIndex(nums []int) int {
	first, second := 0, 1
	if nums[second] > nums[first] {
		first, second = second, first
	}
	for i := 2; i < len(nums); i++ {
		num := nums[i]
		if num > nums[first] {
			second = first
			first = i
		} else if num > nums[second] {
			second = i
		}
	}
	if nums[first] >= 2*nums[second] {
		return first
	}
	return -1
}

func main() {
	fmt.Println(dominantIndex([]int{1, 0}))
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}
