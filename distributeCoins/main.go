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

func distributeCoins(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right)
		ans += abs(left) + abs(right)
		// calculate coins that will move out of this node
		// take all keep one coin
		return node.Val + left + right - 1
	}
	dfs(root)
	// add coins needed from leaves to root
	return ans
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	distributeCoins(root)
}
