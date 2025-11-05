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

func deepestLeavesSum(root *TreeNode) int {
	ans := 0
	maxDepth := findDepth(root)
	findSum(root, 0, maxDepth, &ans)
	return ans
}

func findDepth(node *TreeNode) int {
	if node == nil {
		return -1
	}
	return max(findDepth(node.Left), findDepth(node.Right)) + 1
}

func findSum(node *TreeNode, depth, maxDepth int, ans *int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && depth == maxDepth {
		*ans += node.Val
	}
	findSum(node.Left, depth+1, maxDepth, ans)
	findSum(node.Right, depth+1, maxDepth, ans)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	fmt.Println(deepestLeavesSum(root))
}
