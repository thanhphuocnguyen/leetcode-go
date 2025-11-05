package main

import "fmt"

func maximumXOR(nums []int) int {
	bitFreq := make([]int, 32)
	for _, num := range nums {
		x := num
		j := 0
		for x > 0 {
			if x&1 != 0 {
				bitFreq[j]++
			}
			x >>= 1
			j++
		}
	}
	ans := 0

	for i, freq := range bitFreq {
		if freq != 0 {
			ans |= (1 << i)
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumXOR([]int{1, 2, 3, 9, 2}))
	fmt.Println(maximumXOR([]int{3, 2, 4, 6}))
}
