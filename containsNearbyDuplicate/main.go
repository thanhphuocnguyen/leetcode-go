package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	i, j := 0, 0
	for i <= j && j < len(nums) {
		num := nums[j]
		if j-i <= k {
			mp[num]++
			if mp[num] >= 2 {
				return true
			}
			j++
		} else {
			mp[nums[i]]--
			i++
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
