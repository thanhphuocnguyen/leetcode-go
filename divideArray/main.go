package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)

	i := 0
	for i < n-2 {
		if nums[i+2]-nums[i] <= k {
			ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
			i += 2
		} else {
			return [][]int{}
		}
	}
	return ans
}

func main() {
	// Example usage
	nums := []int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}
	k := 14
	result := divideArray(nums, k)
	for _, triplet := range result {
		println(triplet[0], triplet[1], triplet[2])
	}
}
