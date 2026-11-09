package main

import "fmt"

func maxSum(nums []int, k int) int {
	MOD := 1_000_000_007
	count := make([]int, 32)
	for _, num := range nums {
		x := num
		for i := 0; i < 32; i++ {
			count[i] += (x & 1)
			x >>= 1
		}
	}
	ans := 0
	for i := 0; i < k; i++ {
		x := 0
		for i := 31; i >= 0; i-- {
			x <<= 1
			if count[i] > 0 {
				x |= 1
				count[i]--
			}
		}
		ans += (x * x)
		ans = ans % MOD
	}

	return ans
}

func main() {
	fmt.Println(maxSum([]int{4, 5, 4, 7}, 3))
	fmt.Println(maxSum([]int{2, 6, 5, 8}, 2))
}
