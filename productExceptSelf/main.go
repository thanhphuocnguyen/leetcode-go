package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefixProd := make([]int, n)
	suffixProd := make([]int, n)
	prefixProd[0] = 1
	suffixProd[n-1] = 1
	for i := 1; i < n; i++ {
		prefixProd[i] = prefixProd[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		suffixProd[i] = suffixProd[i+1] * nums[i+1]
	}
	ans := make([]int, n)
	for i := range n {
		ans[i] = prefixProd[i] * suffixProd[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
