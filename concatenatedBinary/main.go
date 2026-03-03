package main

import "fmt"

func concatenatedBinary(n int) int {
	MOD := 1_000_000_007
	shiftCnt := 1
	shiftCheck := 1
	ans := 0
	for i := 1; i <= n; i++ {
		if shiftCheck*2 == i {
			shiftCnt++
			shiftCheck = i
		}
		ans = ((ans << shiftCnt) | i) % MOD
	}
	return ans
}

func main() {
	fmt.Println(concatenatedBinary(3))
}
