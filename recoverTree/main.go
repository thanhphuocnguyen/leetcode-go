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

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	var inOrder func(curr *TreeNode)
	inOrder = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		inOrder(curr.Left)
		if prev != nil && prev.Val > curr.Val {
			if first == nil {
				first = prev
			}
			second = curr
		}
		prev = curr
		inOrder(curr.Right)
	}
	inOrder(root)

	first.Val, second.Val = second.Val, first.Val
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	recoverTree(root)
	fmt.Println(root.Val)
}
