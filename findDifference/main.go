package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	freq := make(map[int]int)
	for _, num := range nums1 {
		freq[num]++
	}
	ans2 := make([]int, 0)
	for _, num := range nums2 {
		if freq[num] == 0 {
			// there are no frequency of num in nums2 in freq of nums1
			// so num is not absent in nums1
			ans2 = append(ans2, num)
			freq[num] = -1
		} else {
			// toogle all intersect to 0
			freq[num] = -1
		}
	}
	ans1 := make([]int, 0)
	for _, num := range nums1 {
		// if f has value so num in nums1 is not absent in nums2
		if freq[num] > 0 {
			ans1 = append(ans1, num)
			freq[num] = -1
		}
	}
	return [][]int{ans1, ans2}
}

func main() {
	fmt.Println(findDifference([]int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10}, []int{-1, -40, -44, 41, 10, -43, 69, 10, 2}))
}
