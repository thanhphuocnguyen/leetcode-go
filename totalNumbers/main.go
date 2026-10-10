package main

import "fmt"

func totalNumbers(digits []int) int {
	seen := make(map[int]bool)
	ans := 0
	for i1, d1 := range digits {
		num := 0
		if d1 == 0 {
			continue
		}
		num = d1
		for i2, d2 := range digits {
			if i2 == i1 {
				continue
			}
			num = num*10 + d2
			for i3, d3 := range digits {
				if d3%2 != 0 || i3 == i2 || i3 == i1 {
					continue
				}

				num = num*10 + d3
				if !seen[num] {
					ans++
					seen[num] = true
				}
				num /= 10
			}
			num /= 10
		}
	}
	return ans
}

func main() {
	fmt.Println(totalNumbers([]int{6, 6, 6}))
	fmt.Println(totalNumbers([]int{1, 2, 3, 4}))
}
