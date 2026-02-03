package main

import "fmt"

func isTrionic(nums []int) bool {
	n := len(nums)
	isAsc := true
	if nums[0] >= nums[1] || nums[n-1] < nums[n-2] {
		return false
	}
	cnt := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			return false
		}
		if isAsc {
			if nums[i+1] < nums[i] {
				isAsc = !isAsc
				cnt++
			}
		} else {
			if nums[i+1] > nums[i] {
				isAsc = !isAsc
				cnt++
			}
		}
	}
	return cnt == 3
}

func main() {
	fmt.Println(isTrionic([]int{1, 3, 5, 4, 2, 6}))
	fmt.Println(isTrionic([]int{2, 1, 3}))
}
