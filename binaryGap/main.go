package main

import "fmt"

func binaryGap(n int) int {
	currGap := 0
	maxGap := 0
	foundOne := false
	for n > 0 {
		if (n & 1) == 1 {
			if !foundOne {
				foundOne = true
			} else {
				maxGap = max(maxGap, currGap+1)
				currGap = 0
			}
		} else if foundOne {
			currGap++
		}
		n >>= 1
	}
	return maxGap
}

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(8))
	fmt.Println(binaryGap(5))
}
