package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	cnt := 1
	prevP := prices[0]
	var ans int64 = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] == prevP-1 {
			cnt++
		} else {
			ans += int64((cnt + 1) * cnt / 2)
			cnt = 1
		}
		prevP = prices[i]
	}
	return ans + int64((cnt+1)*cnt/2)
}

func main() {
	fmt.Println(getDescentPeriods([]int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 4, 3, 10, 9, 8, 7}))
	fmt.Println(getDescentPeriods([]int{1}))
	fmt.Println(getDescentPeriods([]int{8, 6, 7, 7}))
	fmt.Println(getDescentPeriods([]int{3, 2, 1, 4}))
}
