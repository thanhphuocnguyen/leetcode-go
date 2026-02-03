package main

import (
	"fmt"
	"math"
)

func sumFourDivisors(nums []int) int {
	ans := 0
	for _, num := range nums {
		// 1 and itself
		cnt := 0
		sum := 0
		lim := int(math.Sqrt(float64(num)))
		for d := 1; d <= lim; d++ {
			if num%d == 0 {
				q := num / d
				if d == q {
					cnt++
					sum += d
				} else {
					cnt += 2
					sum += d + q
				}
			}
		}
		if cnt == 4 {
			ans += sum
		}
	}
	return ans
}
func main() {
	fmt.Println(sumFourDivisors([]int{21, 4, 7}))
}
