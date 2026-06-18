package main

import (
	"fmt"
	"math"
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

func minDiffInBST(root *TreeNode) int {
	ans := math.MaxInt32
	arr := make([]int, 0)
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		ans = min(ans, arr[i]-arr[i-1])
	}
	return ans
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

func main() {
	root := &TreeNode{
		Val: 90,
		Left: &TreeNode{
			Val: 69,
			Left: &TreeNode{
				Val: 49,
				Right: &TreeNode{
					Val: 52,
				},
			},
			Right: &TreeNode{
				Val: 89,
			},
		},
	}
	fmt.Println(minDiffInBST(root))
}
