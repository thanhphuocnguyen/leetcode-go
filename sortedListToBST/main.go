package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	return buildTree(head, nil)
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	temp := left
	n := 0
	for temp != right {
		n++
		temp = temp.Next
	}

	mid := left
	for i := 0; i < n/2; i++ {
		mid = mid.Next
	}

	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)

	return root
}

func main() {
	list := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{Val: 9},
				},
			},
		},
	}
	tree := sortedListToBST(list)
	fmt.Println(tree)
}
