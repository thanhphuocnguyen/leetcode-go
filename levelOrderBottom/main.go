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

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	var levelOrderTraversal func(node *TreeNode, level int)
	levelOrderTraversal = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(ans) {
			ans = append(ans, []int{})
		}
		levelOrderTraversal(node.Left, level+1)
		ans[level] = append(ans[level], node.Val)
		levelOrderTraversal(node.Right, level+1)
	}
	levelOrderTraversal(root, 0)
	n := len(ans)
	for i := range n / 2 {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}

	return ans
}
