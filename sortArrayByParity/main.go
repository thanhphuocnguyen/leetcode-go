package main

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	i, j := 0, n-1
	for i < j {
		for i < n && (nums[i]&1) == 0 {
			i++
		}

		for j >= 0 && (nums[j]&1) != 0 {
			j--
		}

		if i < j && (nums[i]&1) != 0 && (nums[j]&1) == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}

	}
	return nums
}
func main() {
	sortArrayByParity([]int{0, 1})
}
