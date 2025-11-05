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

func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	dfs(root, "", &ans)
	return ans
}

func dfs(root *TreeNode, path string, ans *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if len(path) != 0 {
			path += "->"
		}
		path += strconv.Itoa(root.Val)
		*ans = append(*ans, path)
		return
	}
	if len(path) != 0 {
		path += "->"
	}
	path += strconv.Itoa(root.Val)
	dfs(root.Left, path, ans)
	dfs(root.Right, path, ans)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
			Left: nil,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(binaryTreePaths(root))
}
