package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	binSearch := func(lo, hi, target int) int {
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if numbers[mid] == target {
				return mid
			} else if numbers[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		return -1
	}
	for i, num := range numbers {
		j := binSearch(i+1, len(numbers)-1, target-num)
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return []int{0, 1}
}

func main() {
	fmt.Println(twoSum([]int{5, 25, 75}, 100))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}
