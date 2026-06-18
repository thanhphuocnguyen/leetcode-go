package main

import "fmt"

func numberOfPoints(nums [][]int) int {
	mini, maxi := nums[0][0], nums[0][1]
	for _, num := range nums {
		mini = min(mini, num[0])
		maxi = max(maxi, num[1])
	}
	n := maxi - mini + 1
	diff := make([]int, n)

	for _, num := range nums {
		l, r := num[0]-mini, num[1]-mini
		diff[l]++
		if r+1 < n {
			diff[r+1]--
		}
	}
	pfxSum := 0
	ans := 0
	for _, d := range diff {
		pfxSum += d
		if pfxSum > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfPoints([][]int{{3, 6}, {1, 5}, {4, 7}}))
}
