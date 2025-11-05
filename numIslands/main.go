package main

import "fmt"

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	// union find
	parents := make([]int, m*n)
	for i := range n * m {
		parents[i] = i
	}
	ranks := make([]int, m*n)
	var find func(num int) int

	find = func(num int) int {
		if parents[num] != num {
			parents[num] = find(parents[num])
		}
		return parents[num]
	}

	union := func(x, y int) {
		xRoot, yRoot := find(x), find(y)
		if xRoot == yRoot {
			return
		}
		if ranks[xRoot] < ranks[yRoot] {
			parents[xRoot] = yRoot
		} else if ranks[yRoot] < ranks[xRoot] {
			parents[yRoot] = xRoot
		} else {
			parents[yRoot] = xRoot
			ranks[xRoot]++
		}
	}

	isValid := func(row, col int) bool {
		return row >= 0 && col >= 0 && row < m && col < n
	}

	getIndex := func(row, col int) int {
		return row*n + col
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i, row := range grid {
		for j, b := range row {
			if b == '1' {
				for _, dir := range dirs {
					nextR, nextC := i+dir[0], j+dir[1]
					if isValid(nextR, nextC) {
						currIdx, nextIdx := getIndex(i, j), getIndex(nextR, nextC)
						union(currIdx, nextIdx)
					}
				}
			}
		}
	}

	mp := make(map[int]bool)
	for i, row := range grid {
		for j, b := range row {
			if b == '1' {
				mp[find(getIndex(i, j))] = true
			}
		}
	}

	return len(mp)
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '1', '1', '1'}}))
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}
