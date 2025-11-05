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

var ans = make([]int, 0)

func largestValues(root *TreeNode) []int {

	dfs(root, 0, ans)
	return ans
}

func dfs(node *TreeNode, level int, list []int) {
	if node == nil {
		return
	}
	if len(list) == level {
		list = append(list, node.Val)
	} else {
		list[level] = max(node.Val, list[level])
	}
	dfs(node.Left, level+1, list)
	dfs(node.Right, level+1, list)
}

func main() {
	// Test data
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(largestValues(root))
}
