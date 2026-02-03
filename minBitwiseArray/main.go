package main

import "fmt"

func main() {
	fmt.Println(minBitwiseArray([]int{2, 3, 5, 7}))
}

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		x := -1
		d := 1
		for (num & d) != 0 {
			x = num - d
			d <<= 1
		}
		ans[i] = x

	}
	return ans
}
