package main

import (
	"fmt"
	"math"
)

func pathInZigZagTree(label int) []int {
	ans := make([]int, 0)
	ans = append(ans, label)
	for label != 1 {
		level := int(math.Log2(float64(label)))
		startAtLevel := int(math.Exp2(float64(level)))
		offset := label - startAtLevel
		offsetParent := offset / 2
		label = label - (offset + offsetParent) - 1
		ans = append(ans, label)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func main() {
	fmt.Println(pathInZigZagTree(26))
	fmt.Println(pathInZigZagTree(14))
}
