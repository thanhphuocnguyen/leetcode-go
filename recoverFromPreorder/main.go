package main

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

func recoverFromPreorder(traversal string) *TreeNode {
	i, n := 0, len(traversal)
	var dfs func(depth int) *TreeNode
	dfs = func(depth int) *TreeNode {
		if i >= n {
			return nil
		}
		currDepth := 0
		for traversal[i] == '-' {
			currDepth++
			i++
		}
		if currDepth != depth {
			i -= currDepth
			return nil
		}
		x := 0

		for traversal[i] != '-' {
			x = x*10 + int(x-'0')
			i++
		}
		node := TreeNode{
			Val: x,
		}
		node.Left = dfs(depth + 1)
		node.Right = dfs(depth + 1)
		return &node
	}
	return dfs(0)
}
