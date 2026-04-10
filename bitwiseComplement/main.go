package main

import "fmt"

func bitwiseComplement(n int) int {
	num := n
	comp := 0
	for num > 0 {
		comp <<= 1
		comp |= 1
		num >>= 1
	}
	return comp - n
}

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
}
