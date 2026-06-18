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

func findTarget(root *TreeNode, k int) bool {
	arr := make([]int, 0)
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		arr = append(arr, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	// two ptr
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i]+arr[j] > k {
			j--
		} else if arr[i]+arr[j] < k {
			i++
		} else {
			return true
		}
	}

	return false
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(findTarget(root, 9))
}
