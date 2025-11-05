package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	for i, j := 1, n-1; i < n; i, j = i+1, j-1 {
		a, b := i, j
		aValid, bValid := true, true
		for a > 0 {
			if a%10 == 0 {
				aValid = false
				break
			}
			a /= 10
		}
		if !aValid {
			continue
		}
		for b > 0 {
			if b%10 == 0 {
				bValid = false
				break
			}
			b /= 10
		}
		if !bValid {
			continue
		}
		return []int{i, j}
	}
	return []int{}
}

func main() {
	fmt.Println(getNoZeroIntegers(1010))
}
