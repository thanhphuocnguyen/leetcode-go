package main

import "fmt"

func maxSumOfSquares(num int, sum int) string {
	tot := 0
	bs := make([]byte, 0, 2000000)
	for i := 0; i < num; i++ {
		for d := 9; d >= 0; d-- {
			if tot+d <= sum {
				bs = append(bs, byte('0'+d))
				tot += d
				break
			}
		}
	}
	if tot < sum {
		return ""
	}
	return string(bs)
}

func main() {
	fmt.Println(maxSumOfSquares(2, 17))
}
