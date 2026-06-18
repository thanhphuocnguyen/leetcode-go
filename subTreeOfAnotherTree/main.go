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

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val {
		return isIdentical(root, subRoot)
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}

func isIdentical(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if (root == nil && subRoot != nil) || (subRoot == nil && root != nil) || root.Val != subRoot.Val {
		return false
	}

	return isIdentical(root.Left, subRoot.Left) && isIdentical(root.Right, subRoot.Right)
}

func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	subRoot := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isSubtree(root, subRoot))
}
