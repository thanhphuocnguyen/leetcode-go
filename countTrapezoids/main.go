package main

import "fmt"

func countTrapezoids(points [][]int) int {
	const MOD int = 1_000_000_007
	mp := make(map[int]int)
	for _, point := range points {
		mp[point[1]]++
	}

	ans := 0
	pairs := 0
	if len(mp) > 1 {
		for _, v := range mp {
			// combination of 2 points in n points
			combine := (v * (v - 1) / 2)
			// sum prev combinations with combination of ways multiply with current points combination
			ans = (ans + (combine * pairs)) % MOD
			// combination between ways prev combinations
			pairs = (pairs + combine) % MOD
		}
	}
	return ans
}

func main() {
	fmt.Println(countTrapezoids([][]int{{-73, -72}, {-1, -56}, {-92, 30}, {-57, -89}, {-19, -89}, {-35, 30}, {64, -72}}))
	fmt.Println(countTrapezoids([][]int{{87, -39}, {12, -94}, {-30, -11}, {-76, -11}}))
	fmt.Println(countTrapezoids([][]int{{50, -41}, {64, -66}, {-45, -41}, {-58, 10}, {25, 10}}))
	fmt.Println(countTrapezoids([][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}))
}
