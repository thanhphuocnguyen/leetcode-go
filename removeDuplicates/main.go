package main

import "fmt"

func removeDuplicates(nums []int) int {
	l, r := 0, 0
	for r < len(nums) {
		for nums[l] == nums[r] {
			r++
		}
		if r-l > 1 && nums[l] != nums[r] {
			nums = append(nums[:l+1], nums[r:]...)
			r = l + 1
		} else {
			l++
			r++
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
