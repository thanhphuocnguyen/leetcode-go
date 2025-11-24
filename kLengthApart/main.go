package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	prevIdx := -1
	for i, num := range nums {
		if num == 0 {
			continue
		} else {
			if prevIdx != -1 && i-prevIdx-1 < k {
				return false
			}
			prevIdx = i
		}
	}
	return true
}

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}
