package main

import (
	"fmt"
	"strconv"
)

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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mp := make(map[string]int)
	var dfs func(root *TreeNode) string
	ans := make([]*TreeNode, 0)
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		serial := strconv.Itoa(root.Val) + "," + left + "," + right
		if mp[serial] == 1 {
			ans = append(ans, root)
		}
		mp[serial]++
		return serial
	}
	dfs(root)
	return ans
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(findDuplicateSubtrees(root))
}
