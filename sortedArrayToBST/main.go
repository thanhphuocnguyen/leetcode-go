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

func sortedArrayToBST(nums []int) *TreeNode {
	mid := (len(nums) - 1) / 2
	root := &TreeNode{
		nums[mid],
		buildTree(nums, 0, mid),
		buildTree(nums, mid, len(nums)-1),
	}
	return root
}

func buildTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	newNode := &TreeNode{nums[mid], buildTree(nums, left, mid-1), buildTree(nums, mid+1, right)}
	return newNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(nums)
}
