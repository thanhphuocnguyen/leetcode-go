package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	prefixProd := 1
	l := 0
	ans := 0
	for r := 0; r < len(nums) && l <= r; r++ {
		prefixProd *= nums[r]
		for l < len(nums) && prefixProd >= k {
			prefixProd /= nums[l]
			l++
		}

		if l <= r {
			ans += r - l + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22}, 18))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
