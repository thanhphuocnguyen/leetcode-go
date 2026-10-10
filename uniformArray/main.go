package main

import (
	"fmt"
	"math"
)

func uniformArray(nums1 []int) bool {
	oddCnt, evenCnt := 0, 0
	minOdd, minEven := math.MaxInt32, math.MaxInt32
	for _, num := range nums1 {
		if num%2 == 0 {
			evenCnt++
			minEven = min(minEven, num)
		} else {
			oddCnt++
			minOdd = min(minOdd, num)
		}
	}

	// case all odd or all even
	if oddCnt == len(nums1) || evenCnt == len(nums1) {
		return true
	}

	ans := true
	// convert to all odd
	for _, num := range nums1 {
		if num%2 == 0 && num-minOdd < 1 {
			ans = false
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(uniformArray([]int{1, 4, 7}))
	fmt.Println(uniformArray([]int{2, 3}))
	fmt.Println(uniformArray([]int{4, 6}))
}
