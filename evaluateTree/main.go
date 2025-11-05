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

func evaluateTree(root *TreeNode) bool {
	return dfs(root, false)
}

func dfs(root *TreeNode, val bool) bool {
	if root == nil {
		return val
	}

	leftVal := dfs(root.Left, val)
	rightVal := dfs(root.Right, val)
	switch root.Val {
	case 2:
		return (leftVal || rightVal)
	case 3:
		return leftVal && rightVal
	case 1:
		return true
	default:
		return false
	}
}

func main() {
	root := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(evaluateTree(root))
}
