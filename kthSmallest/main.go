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

func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	idx := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || idx >= k {
			return
		}
		dfs(node.Left)
		if idx == k-1 {
			ans = node.Val
		}
		idx++
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(kthSmallest(root, 3))
}
