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

func averageOfSubtree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(node *TreeNode, ans *int) (int, int) {
	if node == nil {
		return 0, 0
	}

	sumLeft, cntLeft := dfs(node.Left, ans)
	sumRight, cntRight := dfs(node.Right, ans)
	sum := node.Val + sumLeft + sumRight
	cnt := 1 + cntLeft + cntRight
	avg := sum / cnt
	if node.Val == avg {
		*ans++
	}
	return sum, cnt
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(averageOfSubtree(root))
}
