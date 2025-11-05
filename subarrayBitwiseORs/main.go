package main

import "fmt"

func subarrayBitwiseORs(arr []int) int {
	prev := make(map[int]bool)
	ans := make(map[int]bool)
	for _, num := range arr {
		orMp := make(map[int]bool)
		orMp[num] = true
		for k := range prev {
			orMp[k|num] = true
		}
		for k := range orMp {
			ans[k] = true
		}
		prev = orMp
	}
	return len(ans)
}

func main() {
	fmt.Println(subarrayBitwiseORs([]int{1, 11, 6, 11}))
}
