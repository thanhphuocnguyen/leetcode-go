package main

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	ans := 0
	for _, num := range nums {
		mod := num % k
		for prev := range k {
			dp[prev][mod] = dp[mod][prev] + 1
			ans = max(ans, dp[prev][mod])
		}
	}

	return ans
}

func main() {
	maximumLength([]int{1, 2, 3, 4, 5}, 2)
}
