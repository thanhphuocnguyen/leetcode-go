package main

func sumCounts(nums []int) int {
	ans := 0
	for i, num := range nums {
		mp := make(map[int]bool)
		ans += 1
		mp[num] = true
		for j := i + 1; j < len(nums); j++ {
			mp[nums[j]] = true
			ans += len(mp) * len(mp)
		}
	}
	return ans
}

func main() {
	sumCounts([]int{1, 2, 1})
}
