package main

import "fmt"

func singleNumber(nums []int) int {
	var ans int32 = 0
	for bit := range 32 {
		var cnt int32 = 0
		for _, num := range nums {
			cnt += ((int32(num) >> bit) & 1)
		}
		if cnt%3 != 0 {
			ans |= (1 << bit)
		}
	}
	return int(ans)
}

func main() {
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}
