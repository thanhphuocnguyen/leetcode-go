package main

func countKDifference(nums []int, k int) int {
	counter := make(map[int]bool)
	ans := 0
	for _, num := range nums {
		if counter[num] {
			ans++
		} else {
			counter[num+k] = true
		}
	}
	return ans
}

func main() {
	countKDifference([]int{1, 2, 2, 1}, 1)
}
