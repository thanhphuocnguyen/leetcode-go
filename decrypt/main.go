package main

import "fmt"

func decrypt(code []int, k int) []int {
	sum := 0
	n := len(code)
	left, right := 1, k
	// k < 0 then we will start from n+k -> n-1, ex: k = -2 n = 4 so from [2,3]
	// else will be [1,k]
	if k < 0 {
		k = -k
		left, right = n-k, n-1
	}
	for i := left; i <= right; i++ {
		sum += code[i]
	}
	// sliding window
	ans := make([]int, n)
	for i := range n {
		// value will be current sum
		ans[i] = sum
		// slide and remove left
		sum -= code[left%n]
		left++
		// slide to get sum to the right
		right++
		sum += code[right%n]
	}
	return ans
}

func main() {
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}
