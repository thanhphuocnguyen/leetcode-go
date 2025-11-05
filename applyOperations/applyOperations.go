package main

import "fmt"

func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
			i += 1
		}
	}
	i, j := 0, 1
	for j < n && i < n {
		if nums[i] != 0 {
			i++
			j = i + 1
		} else if nums[j] == 0 {
			j++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		}
	}
	return nums
}

func main() {
	nums := []int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272}
	fmt.Println(applyOperations(nums))
}
