package main

import "fmt"

func sumAndMultiply(n int) int64 {
	sum := 0
	cp := n
	x := 0
	offset := 1
	for cp > 0 {
		d := cp % 10
		if d == 0 {
			cp /= 10
			continue
		}
		sum += d
		x = offset*d + x
		offset *= 10
		cp /= 10
	}
	return int64(x) * int64(sum)
}

func main() {
	fmt.Println(sumAndMultiply(0))
	fmt.Println(sumAndMultiply(10203004))
}
