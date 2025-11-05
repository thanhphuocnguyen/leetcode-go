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

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	ans := 0
	var dfs func(node *TreeNode, maxi, mini int)
	dfs = func(node *TreeNode, maxi, mini int) {
		if node == nil {
			return
		}
		maxi = max(maxi, node.Val)
		mini = min(mini, node.Val)
		dfs(node.Left, maxi, mini)
		dfs(node.Right, maxi, mini)
		ans = max(abs(maxi-mini), ans)
	}
	dfs(root, root.Val, root.Val)
	return ans
}

func main() {
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 7}},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val: 13,
				},
			},
		},
	}
	maxAncestorDiff(root)
}
