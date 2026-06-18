package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMp := make(map[int]*TreeNode)
	children := make(map[int]*TreeNode)
	for _, d := range descriptions {
		parent, child, flag := d[0], d[1], d[2]
		if _, ok := nodeMp[parent]; !ok {
			nodeMp[parent] = &TreeNode{Val: parent}
		}
		if _, ok := nodeMp[child]; !ok {
			nodeMp[child] = &TreeNode{Val: child}
		}
		if flag == 1 {
			nodeMp[parent].Left = nodeMp[child]
		} else {
			nodeMp[parent].Right = nodeMp[child]
		}
		children[child] = nodeMp[parent]
	}
	for k, v := range nodeMp {
		if _, ok := children[k]; !ok {
			return v
		}
	}
	return nil
}

func main() {
	fmt.Println(createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}))
}
