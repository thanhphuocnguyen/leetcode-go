package main

import "fmt"

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
		i--
	}
	if digits[0] == 0 {
		ans := make([]int, len(digits)+1)
		ans[0] = 1
		return ans
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{4, 3, 9, 9}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}
