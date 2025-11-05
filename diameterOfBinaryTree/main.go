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

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	lHeight := dfs(node.Left, ans)
	rHeight := dfs(node.Right, ans)
	*ans = max(*ans, lHeight+rHeight)
	return 1 + max(lHeight, rHeight)
}

func main() {
	diameterOfBinaryTree(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}})
}
