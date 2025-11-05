package main

import "math"

func maxDiff(num int) int {
	// x max y min, a replace max , b replace min
	a, b, x, y := -1, -1, 0, 0
	n := int(math.Log10(float64(num))) + 1
	digits := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		digits[i] = num % 10
		num /= 10
	}

	// find max
	for _, d := range digits {
		if d != 9 {
			a = d
			break
		}
	}
	for _, d := range digits {
		if d == a {
			x = x*10 + 9
		} else {
			x = x*10 + d
		}
	}
	// find min
	if digits[0] != 1 {
		b := digits[0]
		for _, d := range digits {
			if b == d {
				y = y*10 + 1
			} else {
				y = y*10 + d
			}
		}
	} else {
		for i := 1; i < n && b == -1; i++ {
			if digits[i] != 0 && digits[i] != 1 {
				b = digits[i]
			}
		}

		for i, d := range digits {
			if i == 0 {
				y = y*10 + d
			} else if d == b {
				y = y * 10
			} else {
				y = y*10 + d
			}
		}
	}
	return x - y
}

func main() {
	// Example usage
	num := 123456
	result := maxDiff(num)
	println("Max difference:", result) // Output: Max difference: 820000
}
