package main

import "fmt"

func countMajoritySubarrays(nums []int, target int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		freq := map[int]int{}

		for j := i; j < len(nums); j++ {
			freq[nums[j]]++
			if _, ok := freq[target]; !ok {
				continue
			} else if freq[target] > (j-i+1)/2 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(countMajoritySubarrays([]int{1, 2, 2, 3}, 2))
}
