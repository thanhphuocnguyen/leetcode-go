package main

import "fmt"

func predictTheWinner(nums []int) bool {
	n := len(nums)
	// memo := make([]int, n)
	// for i := range n {
	// 	memo[i] = -1
	// }
	var dp func(p1Turn bool, s1, s2 int, i, j int) bool
	dp = func(p1Turn bool, s1, s2 int, i, j int) bool {
		if i > j {
			return s1 >= s2
		}
		// take first
		if p1Turn {
			takeLeft := dp(false, s1+nums[i], s2, i+1, j)
			takeRight := dp(false, s1+nums[j], s2, i, j-1)
			return takeLeft || takeRight

		} else {
			takeLeft := dp(true, s1, s2+nums[i], i+1, j)
			takeRight := dp(true, s1, s2+nums[j], i, j-1)
			return takeLeft && takeRight
		}

	}
	return dp(true, 0, 0, 0, n-1)
}

func main() {
	fmt.Println(predictTheWinner([]int{1, 5, 2, 4, 6}))
	fmt.Println(predictTheWinner([]int{1, 5, 233, 7}))
	fmt.Println(predictTheWinner([]int{1, 5, 2}))
}
