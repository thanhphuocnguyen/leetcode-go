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

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		num = num*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sum += num
			return
		}
		dfs(node.Left, num)
		dfs(node.Right, num)
	}
	dfs(root, 0)
	return sum
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(sumNumbers(root))
}
