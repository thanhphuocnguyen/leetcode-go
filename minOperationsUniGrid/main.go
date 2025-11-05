package main

import "sort"

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	flatten := make([]int, m*n)
	for i, cell := range grid {
		for j, num := range cell {
			flatten[i*n+j] = num
		}
	}
	sort.Ints(flatten)

	median := flatten[(0+len(grid)-1)/1]
	ans := 0
	for _, num := range flatten {
		quotient := 0
		if num > median {
			quotient = num - median
		} else {
			quotient = median - num
		}
		if quotient%x != 0 {
			return -1
		}
		ans += quotient / x
	}
	return ans
}

func main() {
	println(minOperations([][]int{{980, 476, 644, 56}, {644, 140, 812, 308}, {812, 812, 896, 560}, {728, 476, 56, 812}}, 84))

}
