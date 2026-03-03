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

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	inOrderTraverse(root, &nums)
	return constructNewRoot(nums, 0, len(nums)-1)
}

func inOrderTraverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrderTraverse(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrderTraverse(root.Right, nums)
}

func constructNewRoot(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = constructNewRoot(nums, left, mid-1)
	root.Right = constructNewRoot(nums, mid+1, right)
	return root
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(balanceBST(root))
}
