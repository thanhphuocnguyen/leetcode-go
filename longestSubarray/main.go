package main

import "fmt"

func longestSubarray(nums []int) int {
	i := 0
	for nums[i] == 0 {
		i++
	}
	j := i
	ans, currCnt, zeroCnt := 0, 0, 0
	for j < len(nums) {
		if zeroCnt > 1 {
			if nums[j] == 0 {
				currCnt = 0
				j++
				continue
			} else {
				currCnt = 1
				zeroCnt = 0
			}
		}
		if nums[j] == 1 {
			currCnt++
			ans = max(ans, currCnt)
		} else {
			zeroCnt++
		}
		j++
	}
	return ans
}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}
