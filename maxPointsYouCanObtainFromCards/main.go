package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	// we should be reverse thinking in this case
	// we should remove n-k window points from the total sum to calculate the largest points after left right pick times
	prefixSum := make([]int, n+1)
	prefixSum[0] = 0
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + cardPoints[i-1]
	}
	i, j := 0, n-k
	ans := 0
	for j <= n {
		windowSum := prefixSum[j] - prefixSum[i]
		ans = max(ans, prefixSum[n]-windowSum)
		i++
		j++
	}

	return ans
}

func main() {
	fmt.Println(maxScore([]int{96, 90, 41, 82, 39, 74, 64, 50, 30}, 8))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}
