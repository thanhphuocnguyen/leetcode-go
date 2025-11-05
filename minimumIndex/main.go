package main

func minimumIndex(nums []int) int {
	n := len(nums)
	freqMap := make(map[int]int)
	dominant := nums[0]
	for _, num := range nums {
		freqMap[num]++
		if freqMap[num] > freqMap[dominant] {
			dominant = num
		}
	}
	f1 := 0
	i := 0
	for i <= n-2 {
		if nums[i] == dominant {
			f1++
			if f1*2 > i+1 {
				break
			}
		}
		i++
	}
	f2 := freqMap[dominant] - f1
	if f2*2 > n-1-i && f1*2 > i {
		return i
	}
	return -1
}

func main() {
	println(minimumIndex([]int{1}))
	println(minimumIndex([]int{1, 2, 2}))
	println(minimumIndex([]int{1, 2, 2, 2}))
	println(minimumIndex([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
	println(minimumIndex([]int{3, 3, 3, 3, 7, 2, 2}))
}
