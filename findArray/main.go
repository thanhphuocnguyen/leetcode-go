package main

import "fmt"

func findArray(pref []int) []int {
	ans := make([]int, len(pref))
	ans[0] = pref[0]
	xor := pref[0]
	for i := 1; i < len(pref); i++ {
		ans[i] = xor ^ pref[i]
		xor ^= ans[i]
	}
	return ans
}

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1}))
}
