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

func goodNodes(root *TreeNode) int {
	// root is always counted
	ans := 0
	var dfs func(node *TreeNode, currMax int)

	dfs = func(node *TreeNode, currMax int) {
		if node == nil {
			return
		}
		if node.Val >= currMax {
			ans++
			currMax = max(currMax, node.Val)
		}
		dfs(node.Left, currMax)
		dfs(node.Right, currMax)
	}

	dfs(root, root.Val)
	return ans
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 5}

	println(goodNodes(root))
}
