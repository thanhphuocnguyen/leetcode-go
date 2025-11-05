package main

import "fmt"

func getSneakyNumbers(nums []int) []int {
	var ans = []int{0, 0}
	var mask int64
	i := 0
	for _, num := range nums {
		if ((1 << num) & mask) != 0 {
			ans[i] = num
			i++
		} else {
			mask = mask | (1 << num)
		}
	}
	return ans
}

func main() {
	fmt.Println((getSneakyNumbers([]int{43, 29, 4, 51, 10, 31, 64, 65, 48, 23, 21, 13, 5, 66, 24, 46, 49, 33, 47, 41, 38, 39, 35, 30, 27, 40, 9, 12, 20, 62, 58, 63, 11, 22, 42, 32, 60, 44, 14, 67, 16, 37, 0, 56, 2, 45, 53, 1, 61, 3, 55, 25, 52, 17, 36, 6, 28, 34, 26, 57, 43, 15, 54, 18, 68, 59, 19, 66, 50, 7, 8})))
}
