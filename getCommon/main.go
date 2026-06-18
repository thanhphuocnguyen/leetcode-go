package main

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if nums1[n1-1] < nums2[0] || nums2[n2-1] < nums1[0] {
		return -1
	}

	i1, i2 := 0, 0
	for i1 < n1 && i2 < n2 {
		if nums1[i1] < nums2[i2] {
			i1++
		} else if nums1[i1] > nums2[i2] {
			i2++
		} else {
			return nums1[i1]
		}
	}
	return -1
}

func main() {
	fmt.Println(getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}))
}
