package main

import "fmt"

func resultArray(nums []int, k int) []int64 {
	n := len(nums)
	dp := make([]int64, k)
	// ans is find number of prefix, suffix sub-arrays that divide by k remain i (0,1,2,3....)
	ans := make([]int64, k)

	// extend sub-array
	for i := 0; i < n; i++ {
		newDp := make([]int64, k)
		// start new sub-array at nums[i]
		rem := nums[i] % k
		// init state will be 1
		newDp[rem] += 1
		// compute for every rem for current sub array
		for r := range k {
			newRem := (r * nums[i]) % k
			// state of new rem will include previous rem state
			newDp[newRem] += dp[r]
		}
		/*
			[1]
			0, 1, 2
			-> 0, 1, 2
			-> newDp[0] += dp[0], newDp[1]+=dp[1], newDp[2] += dp[2]
		*/
		for r := range k {
			ans[r] += newDp[r]
		}
		dp = newDp
	}

	return ans
}

func main() {
	fmt.Println(resultArray([]int{1, 2, 3, 4, 5}, 3))
}
