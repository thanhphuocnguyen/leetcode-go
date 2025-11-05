package main

import "fmt"

func totalMoney(n int) int {
	circle := n / 7
	remainder := n % 7
	cnt := 0
	ans := 0
	start := 7
	for circle > 0 {
		ans += start*(start+1)/2 - (cnt * (cnt + 1) / 2)
		start++
		cnt++
		circle--
	}
	ans += (remainder+cnt)*(remainder+cnt+1)/2 - (cnt * (cnt + 1) / 2)
	return ans
}

func main() {
	fmt.Println(totalMoney(4))
	fmt.Println(totalMoney(10))
	fmt.Println(totalMoney(20))
}
