package main

import (
	"fmt"
	"math"
)

func isThree(n int) bool {
	root := int(math.Sqrt(float64(n)))
	if root*root != n {
		return false
	}

	if root < 2 {
		return false
	}

	for i := 2; i*i <= root; i++ {
		if root%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isThree(4))
}
