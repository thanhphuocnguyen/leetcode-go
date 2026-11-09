package main

import "fmt"

func sortPermutation(nums []int) int {
	n := len(nums)
	k := -1
	// maximum k for every swap is eq AND of all elements that not in order
	for i := 0; i < n; i++ {
		if nums[i] != i {
			if k == -1 {
				k |= nums[i]
			} else {
				k &= nums[i]
			}
		}
	}

	if k == -1 {
		return 0
	}
	return k
}

func main() {
	fmt.Println(sortPermutation([]int{0, 3, 2, 1}))
	fmt.Println(sortPermutation([]int{0, 3, 2, 1}))
}
