package main

import "fmt"

func findMatrix(nums []int) [][]int {
	freq := make([]int, len(nums)+1)
	ans := make([][]int, 0)
	for _, num := range nums {
		if freq[num] > len(ans) {
			ans = append(ans, []int{})
		}
		ans[freq[num]] = append(ans[freq[num]], num)
		freq[num]++
	}
	return ans
}

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
}
