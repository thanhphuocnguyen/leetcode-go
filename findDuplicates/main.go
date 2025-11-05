package main

func findDuplicates(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		idx := abs(num) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			ans = append(ans, abs(num))
		}
	}

	return ans
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
}
