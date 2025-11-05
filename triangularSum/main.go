package main

func triangularSum(nums []int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			nums[j] = (nums[j] + nums[j+1]) % 10
		}
	}
	return nums[0]
}

func main() {
	triangularSum([]int{1, 2, 3, 4, 5})
}
