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

func inorderTraversal(root *TreeNode) []int {
	st := make([]*TreeNode, 0)
	ans := make([]int, 0)
	currNode := root
	for currNode != nil || len(st) > 0 {
		for currNode != nil {
			st = append(st, currNode)
			currNode = currNode.Left
		}
		currNode := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, currNode.Val)
		currNode = currNode.Right
	}
	return ans
}

func main() {
	// Test cases
	// [1, null, 2, 3]
	// Output: [1, 3, 2]
	// Explanation: The root node is 1, the right child node of root node is 2, and the right child node of 2 is 3.
	// The inorder traversal is [1, 3, 2].
	// root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	// [1, 2, 3, 4, 5, 6, 7]
	// Output: [4, 2, 5, 1, 6, 3, 7]
	// Explanation: The root node is 1, the left child node of 1 is 2, the right child node of 1 is 3.
	// The left child node of 2 is 4, the right child node of 2 is 5.
	// The left child node of 3 is 6, the right child node of 3 is 7.
	// The inorder traversal is [4, 2, 5, 1, 6, 3, 7].
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	fmt.Println(inorderTraversal(root))
}
