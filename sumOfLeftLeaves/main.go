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

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	dfs(root, &sum)
	return sum
}

func dfs(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*sum += node.Left.Val
	}
	dfs(node.Left, sum)
	dfs(node.Right, sum)
}

func main() {
	root := &TreeNode{
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(sumOfLeftLeaves(root))
}
