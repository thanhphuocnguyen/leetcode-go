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

func findMin(node *TreeNode) *TreeNode {
	temp := node
	rs := temp

	for temp != nil {
		if temp.Val < rs.Val {
			rs = temp
		}
		if temp.Left != nil {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}
	return rs
}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		successor := findMin(root.Right)
		root.Val = successor.Val
		root.Right = deleteNode(root.Right, successor.Val)
	}
	return root
}

func main() {

	fmt.Println()
}
