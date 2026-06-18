package main

import (
	"fmt"
	"math/rand/v2"
)

func findKthLargest(nums []int, k int) int {
	var solve func(l, r int) int
	solve = func(l, r int) int {
		pick := rand.IntN(r - l + 1)
		pivot := pick + l
		i, j := l, r
		for i <= j {
			for nums[i] < nums[pivot] {
				i++
			}
			for nums[j] > nums[pivot] {
				j--
			}

			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if j >= k {
			return solve(l, j)
		}
		if i <= k {
			return solve(i, r)
		}

		return nums[len(nums)-k]
	}

	return solve(0, len(nums)-1)
}
func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
