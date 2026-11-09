package main

var MOD = 1_000_000_007

func numOfWays(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 1
	}

	root := nums[0]
	left := make([]int, 0, n)
	right := make([]int, 0, n)

	leftWays := numOfWays(left)
	rightWays := numOfWays(right)

	for i := 1; i < n; i++ {
		if nums[i] < root {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	interleaving := C(n-1, len(left))

	return interleaving * leftWays * rightWays % MOD
}

func C(n, k int) int {

}
