package main

func highestPeak(isWater [][]int) [][]int {
	heightMap := make([][]int, len(isWater))
	for i := range isWater {
		heightMap[i] = make([]int, len(isWater[i]))
		for j := range isWater[i] {
			if isWater[i][j] == 1 {
				heightMap[i][j] = 0
			} else {
				heightMap[i][j] = -1
			}
		}
	}
	directions := [4][2]int{
		{0, 1},  // r
		{0, -1}, // l
		{1, 0},  // d
		{-1, 0}, // u
	}
}
