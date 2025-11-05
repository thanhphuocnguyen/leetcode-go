package main

import "fmt"

func differenceOfSums(n int, m int) int {
	first := m
	last := findLargeNum(n, m)
	if last > n {
		first -= m
	}
	cnt := n / m
	sumDivisible := cnt * (first + last) / 2
	return n*(n+1)/2 - 2*(sumDivisible)
}

func findLargeNum(n int, m int) int {
	rem := (n + m) % m
	if rem == 0 {
		return n
	} else {
		return n + m - rem
	}
}

func main() {
	n := 10
	m := 3
	result := differenceOfSums(n, m)
	fmt.Println("Result:", result)
}
