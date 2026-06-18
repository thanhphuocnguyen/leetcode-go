package main

import "fmt"

func distance(nums []int) []int64 {
	mp := make(map[int][]int)
	for i, num := range nums {
		if _, ok := mp[num]; !ok {
			mp[num] = []int{i}
		} else {
			mp[num] = append(mp[num], i)
		}
	}

	ans := make([]int64, len(nums))
	for _, indices := range mp {
		total := 0
		for _, i := range indices {
			total += i
		}
		n := len(indices)
		prfSum := 0
		for i, idx := range indices {
			left := int64(i*idx - prfSum)
			right := int64((total - prfSum - idx) - (n-i-1)*idx)
			ans[idx] = left + right
			prfSum += idx
		}
	}
	return ans
}

func main() {
	fmt.Println(distance([]int{1, 3, 1, 1, 2}))
}
