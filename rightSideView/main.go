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

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)

	var dfs func(node *TreeNode, lev int)
	dfs = func(node *TreeNode, lev int) {
		if node == nil {
			return
		}
		if lev == len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, lev+1)
		dfs(node.Left, lev+1)
	}

	dfs(root, 0)
	return ans
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(rightSideView(root))
}
