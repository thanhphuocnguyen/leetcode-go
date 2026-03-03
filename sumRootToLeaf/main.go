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

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, currSum int)
	dfs = func(node *TreeNode, currSum int) {
		tempSum := (currSum << 1) | node.Val
		if node.Left == nil && node.Right == nil {
			sum += tempSum
		}
		if node.Left != nil {
			dfs(node.Left, tempSum)
		}
		if node.Right != nil {
			dfs(node.Right, tempSum)
		}
	}
	dfs(root, 0)
	return sum
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(sumRootToLeaf(root))
}
