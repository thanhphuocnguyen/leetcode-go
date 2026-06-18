package main

import "fmt"

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	layers := min(m, n) / 2
	for layer := range layers {
		vals := make([]int, 0)
		rowIdx := make([]int, 0)
		colIdx := make([]int, 0)
		// rotate left bottom right top
		// left
		for i := layer; i < m-layer-1; i++ {
			rowIdx = append(rowIdx, i)
			colIdx = append(colIdx, layer)
			vals = append(vals, grid[i][layer])
		}
		// down
		for i := layer; i < n-layer-1; i++ {
			rowIdx = append(rowIdx, m-layer-1)
			colIdx = append(colIdx, i)
			vals = append(vals, grid[m-layer-1][i])
		}

		// right
		for i := m - layer - 1; i > layer; i-- {
			rowIdx = append(rowIdx, i)
			colIdx = append(colIdx, n-layer-1)
			vals = append(vals, grid[i][n-layer-1])
		}

		// top
		for i := n - layer - 1; i > layer; i-- {
			rowIdx = append(rowIdx, layer)
			colIdx = append(colIdx, i)
			vals = append(vals, grid[layer][i])
		}
		total := len(vals)
		kmod := k % total
		for i := range total {
			idx := (total + i - kmod) % total
			grid[rowIdx[i]][colIdx[i]] = vals[idx]
		}
	}

	return grid
}
func main() {
	fmt.Println(rotateGrid([][]int{{40, 10}, {30, 20}}, 1))
}
