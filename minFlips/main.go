package main

import "fmt"

func minFlips(grid [][]int) int {
	ans := 0
	flip := func(i, j int) {
		if grid[i][j] == 0 {
			grid[i][j] = 1
		} else {
			grid[i][j] = 0
		}
	}
	for i, row := range grid {
		for j, cell := range row {
			if i-2 >= 0 && cell != grid[i-2][j] {
				ans++
				flip(i, j)
			}
			if j-2 >= 0 && cell != grid[i][j-2] {
				ans++
				flip(i, j)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}
