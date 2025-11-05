package main

func reorderedPowerOf2(n int) bool {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	used := make([]bool, len(nums))
	var backtrack func(temp []int) bool

	backtrack = func(temp []int) bool {
		if temp[0] == 0 {
			return false
		}
		if len(temp) == len(nums) {
			num := getNum(temp)
			return num&(num-1) == 0
		}
		for i, num := range nums {
			if i > 0 && num == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				temp = append(temp, num)
				used[i] = true
				if backtrack(temp) {
					return true
				}
				temp = temp[:len(temp)-1]
				used[i] = false
			}

		}
		return false
	}

	return backtrack([]int{})
}

func getNum(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans = ans*10 + num
	}
	return ans
}
