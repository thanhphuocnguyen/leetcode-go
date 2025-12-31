package main

import "fmt"

var DIRECTIONS = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func latestDayToCross(row int, col int, cells [][]int) int {
	n := len(cells)
	left, right := 1, n
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if possible(row, col, mid, cells) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func possible(row, col, day int, cells [][]int) bool {
	// initialize grid land cells
	visited := make([][]bool, row)
	grid := make([][]int, row)
	for i := range row {
		grid[i] = make([]int, col)
		visited[i] = make([]bool, col)
	}
	// flood land cell til day
	for i := range day {
		cell := cells[i]
		r, c := cell[0]-1, cell[1]-1
		grid[r][c] = 1
	}
	// bfs
	// queue save row, col of current cell
	q := make([][2]int, 0)
	for i := range col {
		if grid[0][i] == 0 {
			q = append(q, [2]int{0, i})
		}
	}

	for len(q) > 0 {
		cell := q[0]
		q = q[1:]
		r, c := cell[0], cell[1]
		if r == row-1 {
			return true
		}
		for _, d := range DIRECTIONS {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < row && nc >= 0 && nc < col && !visited[nr][nc] && grid[nr][nc] == 0 {
				visited[nr][nc] = true
				q = append(q, [2]int{nr, nc})
			}
		}
	}

	return false
}

func main() {
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}))
}
