package main

import "fmt"

func countPartitions(nums []int) int {
	prfSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prfSum[i] = nums[i-1] + prfSum[i-1]
	}
	ans := 0
	for i := 1; i < len(nums); i++ {
		if (prfSum[i]-(prfSum[len(nums)]-prfSum[i]))%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}))
}
