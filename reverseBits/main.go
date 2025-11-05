package main

import "fmt"

func reverseBits(n int) int {
	ans := 0
	for i := 31; i >= 0; i-- {
		if (n & 1) != 0 {
			ans += pow2(i)
		}
		n >>= 1
	}
	return ans
}

func pow2(e int) int {
	ans := 1
	for e > 0 {
		ans *= 2
		e--
	}
	return ans
}

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(2147483644))
}
