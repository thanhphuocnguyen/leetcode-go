package main

import (
	"fmt"
	"sort"
)

func findScore(nums []int) int64 {
	idxArr := make([]int, len(nums))
	for i := range idxArr {
		idxArr[i] = i
	}
	sort.Slice(idxArr, func(i, j int) bool {
		if nums[i] == nums[j] {
			return i < j
		}
		return nums[i] < nums[j]
	})
	var ans int64
	for _, idx := range idxArr {
		if nums[idx] == -1 {
			continue
		}
		ans += int64(nums[idx])
		nums[idx] = -1
		if idx > 0 {
			nums[idx-1] = -1
		}
		if idx < len(nums)-1 {
			nums[idx+1] = -1
		}
	}
	return ans
}

func main() {
	nums := []int{10, 44, 10, 8, 48, 30, 17, 38, 41, 27, 16, 33, 45, 45, 34, 30, 22, 3, 42, 42}
	fmt.Println(findScore(nums))
}
