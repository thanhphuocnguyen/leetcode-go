package main

import "fmt"

func rearrangeArray(nums []int) []int {
	freq := [101]int{}
	n := len(nums)
	for _, num := range nums {
		freq[num]++
	}
	ans := make([]int, 0, 101)
	for i := 0; i < n; i++ {
		for i, f := range freq {
			if f > 0 {
				ans = append(ans, i)
				freq[i]--
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, 3, 2, 1, 3}))
}
