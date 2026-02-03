package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := strconv.Itoa(root.Val)
	if root.Left != nil {
		ans += "(" + tree2str(root.Left) + ")"
	}
	if root.Right != nil {
		if root.Left == nil {
			ans += "()"
		}
		ans += "(" + tree2str(root.Right) + ")"
	}
	return ans
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree2str(root)
}
