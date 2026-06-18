package main

import (
	"fmt"
	"sort"
)

func largestInteger(num int) int {
	cp := num
	odd := make([]int, 0)
	even := make([]int, 0)
	nums := make([]int, 0)
	for cp > 0 {
		d := cp % 10
		if (d & 1) == 1 {
			odd = append(odd, d)
		} else {
			even = append(even, d)
		}
		nums = append(nums, d)
		cp /= 10
	}

	ans := 0

	sort.Ints(odd)
	sort.Ints(even)
	i, j := len(even)-1, len(odd)-1

	for k := len(nums) - 1; k >= 0; k-- {
		d := nums[k]
		if (d & 1) == 1 {
			ans = ans*10 + odd[j]
			j--
		} else {
			ans = ans*10 + even[i]
			i--
		}
	}

	return ans
}
func main() {
	fmt.Println(largestInteger(65875))
	fmt.Println(largestInteger(60))
	fmt.Println(largestInteger(1234))
	fmt.Println(largestInteger(247))
}
