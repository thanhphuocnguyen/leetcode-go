package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum / 2
	memo := make([][]int, len(nums))
	for i := range len(nums) {
		memo[i] = make([]int, target+1)
		for j := range target + 1 {
			memo[i][j] = -1
		}
	}
	return backtrack(0, target, nums, memo)
}

func backtrack(index, currSum int, nums []int, memo [][]int) bool {
	if currSum == 0 {
		return true
	}

	if index >= len(nums) || currSum < 0 {
		return false
	}

	if memo[index][currSum] != -1 {
		return memo[index][currSum] == 1
	}
	skip := backtrack(index+1, currSum, nums, memo)
	take := false
	if nums[index] <= currSum {
		take = backtrack(index+1, currSum-nums[index], nums, memo)
	}
	if skip || take {
		memo[index][currSum] = 1
	} else {
		memo[index][currSum] = 0
	}
	return skip || take
}

func main() {
	nums := []int{1, 2, 3, 5}
	result := canPartition(nums)
	println(result) // Output: true
}
