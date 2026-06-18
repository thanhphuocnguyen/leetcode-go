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

func isCousins(root *TreeNode, x int, y int) bool {
	var px, py *TreeNode
	var dx, dy = 0, 0
	var dfs func(node *TreeNode, pNode *TreeNode, d int)
	dfs = func(node, pNode *TreeNode, d int) {
		if node == nil {
			return
		}
		if node.Val == x {
			px = pNode
			dx = d
			return
		}
		if node.Val == y {
			py = pNode
			dy = d
			return
		}
		dfs(node.Left, node, d+1)
		dfs(node.Right, node, d+1)
	}
	dfs(root, nil, 0)
	if dx == dy && px != py {
		return true
	}
	return false
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(isCousins(root, 5, 4))
}
