package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if len(nums) < 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < n {
		if slow > 0 {
			if nums[fast] == nums[slow-2] {
				fast++
			} else {
				nums[slow] = nums[fast]
				slow++
				fast++
			}
		}
	}
	return slow
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}
