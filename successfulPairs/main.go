package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n := len(potions)
	ans := make([]int, len(spells))
	for i, sp := range spells {
		idx := binarySearch(potions, sp, success)
		if idx < n {
			ans[i] = n - idx
		}
	}
	return ans
}
func binarySearch(arr []int, spell int, success int64) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		var product int64 = int64(arr[mid]) * int64(spell)
		if product < success {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}
