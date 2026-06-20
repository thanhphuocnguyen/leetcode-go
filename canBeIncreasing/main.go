package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	currMin := nums[n-1]
	cntRight := 0
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= currMin {
			cntRight++
		} else {
			currMin = min(nums[i], currMin)
		}
	}
	cntLeft := 0
	currMax := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] <= currMax {
			cntLeft++
		} else {
			currMax = max(nums[i], currMax)
		}
	}

	return cntRight <= 1 || cntLeft <= 1
}

func main() {
	fmt.Println(canBeIncreasing([]int{541, 783, 433, 744}))
	fmt.Println(canBeIncreasing([]int{105, 924, 32, 968}))
	fmt.Println(canBeIncreasing([]int{100, 21, 100}))
	fmt.Println(canBeIncreasing([]int{1, 1, 1}))
	fmt.Println(canBeIncreasing([]int{1, 2, 10, 5, 7}))
	fmt.Println(canBeIncreasing([]int{2, 3, 1, 2}))
}
