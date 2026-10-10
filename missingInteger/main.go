package main

import "fmt"

func missingInteger(nums []int) int {
	mp := make(map[int]bool)
	prefixSum := nums[0]
	mp[nums[0]] = true
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			prefixSum += nums[i]
		} else {
			break
		}
	}
	for ; i < len(nums); i++ {
		mp[nums[i]] = true
	}
	for mp[prefixSum] {
		prefixSum++
	}

	return prefixSum
}

func main() {
	fmt.Println(missingInteger([]int{1, 2, 3, 2, 5}))
	fmt.Println(missingInteger([]int{3, 4, 5, 1, 12, 14, 13}))
}
