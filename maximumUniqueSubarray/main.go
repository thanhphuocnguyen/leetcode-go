package main

import "fmt"

func maximumUniqueSubarray(nums []int) int {
	maxScore, ans := 0, 0
	l := 0
	freq := make(map[int]bool)
	for _, num := range nums {
		if freq[num] {
			for freq[num] {
				maxScore -= nums[l]
				freq[nums[l]] = false
				l++
			}
		}
		maxScore += num
		freq[num] = true
		ans = max(maxScore, ans)
	}
	return ans
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
}
