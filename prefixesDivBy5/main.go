package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	var curr int64 = 0
	ans := make([]bool, n)
	for i, num := range nums {
		curr <<= 1
		curr |= int64(num)
		curr %= 5
		ans[i] = curr == 0
	}

	return ans
}

func main() {
	fmt.Println(prefixesDivBy5([]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}))
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 0, 1}))
}
