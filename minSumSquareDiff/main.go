package main

import (
	"fmt"
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	k := k1 + k2

	diffArr := make([]int, n+1)
	for i := range n {
		diffArr[i] = abs(nums1[i] - nums2[i])
	}
	sort.Slice(diffArr, func(i, j int) bool {
		return diffArr[i] > diffArr[j]
	})

	for i := range n {
		gap := diffArr[i] - diffArr[i+1]

		cnt := i + 1
		cost := gap * cnt
		if k >= cost {
			k -= cost
		} else {
			var ans int64 = 0
			level := diffArr[i] - k/cnt
			rem := k % cnt
			for j := range i + 1 {
				if j < rem {
					ans += int64((level - 1) * (level - 1))
				} else {
					ans += int64(level * level)
				}
			}

			for j := i + 1; j < n; j++ {
				ans += int64(diffArr[j] * diffArr[j])
			}
			return ans
		}

	}

	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(minSumSquareDiff([]int{11, 12, 13, 14, 15}, []int{13, 16, 16, 12, 14}, 3, 6))
	// fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 10, 2))
	// fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
}
