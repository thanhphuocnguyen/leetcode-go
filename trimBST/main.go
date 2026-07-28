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

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Right, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func main() {
	tree1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(trimBST(tree1, 2, 4))

	tree2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(trimBST(tree2, 2, 4))
}
