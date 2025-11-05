package main

import (
	"fmt"
)

func largestGoodInteger(num string) string {
	cnt := 1
	var currMax byte
	ans := ""
	for i := 1; i < len(num); i++ {
		if num[i] == num[i-1] {
			cnt++
			if cnt == 3 && currMax < num[i] {
				currMax = num[i]
				ans = num[i-2 : i+1]
			}
		} else {
			cnt = 1
		}
	}
	return ans
}

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}
