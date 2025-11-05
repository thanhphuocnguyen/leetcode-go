package main

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

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return buildTree(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func buildTree(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if preStart > preEnd || postStart > postEnd {
		return nil
	}
	val := preorder[preStart]
	node := TreeNode{val, nil, nil}
	nextVal := preorder[preStart+1]
	lenLeft := 0
	for i := postStart; i <= postEnd; i++ {
		if postorder[i] == nextVal {
			lenLeft = i - postStart + 1
			break
		}
	}
	node.Left = buildTree(
		preorder, preStart+1, preStart+lenLeft,
		postorder, postStart, postStart+lenLeft-1)
	node.Right = buildTree(
		preorder, preStart+lenLeft+1, preEnd,
		postorder, postStart+lenLeft, postEnd-1)
	return &node
}

func main() {
	constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
}
