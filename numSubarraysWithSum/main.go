package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
	prfSum := 0
	ans := 0
	mp := make(map[int]int)
	mp[prfSum] = 1
	for _, num := range nums {
		prfSum += num
		if mp[prfSum-goal] > 0 {
			ans += mp[prfSum-goal]
		}
		mp[prfSum]++
	}

	return ans
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
}
