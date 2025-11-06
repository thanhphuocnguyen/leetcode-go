package main

type NeighborSum struct {
	n      int
	coords [][2]int
	grid   [][]int
}

func Constructor(grid [][]int) NeighborSum {
	n := len(grid)
	coord := make([][2]int, n*n)
	for i := range n {
		coord[i] = [2]int{}
	}
	for i, row := range grid {
		for j, cell := range row {
			coord[cell] = [2]int{i, j}
		}
	}
	return NeighborSum{n, coord, grid}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	row, col := this.coords[value][0], this.coords[value][1]
	ans := 0
	if isValid(row-1, col, this.n) {
		ans += this.grid[row-1][col]
	}
	if isValid(row, col-1, this.n) {
		ans += this.grid[row][col-1]
	}
	if isValid(row+1, col, this.n) {
		ans += this.grid[row+1][col]
	}
	if isValid(row, col+1, this.n) {
		ans += this.grid[row][col+1]
	}
	return ans
}

func (this *NeighborSum) DiagonalSum(value int) int {
	row, col := this.coords[value][0], this.coords[value][1]
	ans := 0
	if isValid(row-1, col-1, this.n) {
		ans += this.grid[row-1][col-1]
	}
	if isValid(row-1, col+1, this.n) {
		ans += this.grid[row-1][col+1]
	}
	if isValid(row+1, col-1, this.n) {
		ans += this.grid[row+1][col-1]
	}
	if isValid(row+1, col+1, this.n) {
		ans += this.grid[row+1][col+1]
	}
	return ans
}

func isValid(i, j, n int) bool {
	return i >= 0 && j >= 0 && i < n && j < n
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */

func main() {
	ctor := Constructor([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}})
	ctor.AdjacentSum(4)
}
