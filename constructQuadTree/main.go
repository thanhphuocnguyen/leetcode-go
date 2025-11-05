package main

import "fmt"

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	root := getNode(grid, 0, 0, n, n, n)
	return root
}

func getNode(grid [][]int, row, col, toRow, toCol, n int) *Node {
	root := &Node{
		IsLeaf: false,
		Val:    false,
	}

	check := true
	checkCell := grid[row][col]
	for i := row; i < toRow; i++ {
		if !check {
			break
		}
		for j := col; j < toCol; j++ {
			cell := grid[i][j]
			if cell != checkCell {
				check = false
				break
			}
		}
	}
	if check {
		root.Val = checkCell == 1
		root.IsLeaf = true
	} else {
		root.TopLeft = getNode(grid, row, col, row+n/2, col+n/2, n/2)
		root.BottomLeft = getNode(grid, row+n/2, col, row+n, col+n/2, n/2)
		root.TopRight = getNode(grid, row, col+n/2, row+n/2, col+n, n/2)
		root.BottomRight = getNode(grid, row+n/2, col+n/2, row+n, col+n, n/2)
	}
	return root
}

func main() {
	fmt.Println(construct([][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}))
	fmt.Println(construct([][]int{{0, 1}, {1, 0}}))
}
