package main

import "fmt"

func isZeroArray(nums []int, queries [][]int) bool {
	diffArr := buildDifferenceArray(nums)
	for _, q := range queries {
		update(diffArr, q[0], q[1], -1)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[i] = diffArr[0]
		} else {
			nums[i] = nums[i-1] + diffArr[i]
		}
		if nums[i] > 0 {
			return false
		}
	}

	return true
}

func buildDifferenceArray(nums []int) []int {
	diffArr := make([]int, len(nums)+1)
	diffArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffArr[i] = nums[i] - nums[i-1]
	}
	return diffArr
}

func update(arr []int, l, r, v int) {
	arr[l] += v
	arr[r+1] -= v
}

func main() {
	fmt.Println(isZeroArray([]int{7}, [][]int{{0, 0}}))
}
