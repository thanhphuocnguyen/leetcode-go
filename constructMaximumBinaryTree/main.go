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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums))
}

func build(nums []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}
	maxNum, maxIdx := nums[start], start
	for i := start + 1; i < end; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIdx = i
		}
	}
	root := &TreeNode{
		Val: maxNum,
	}
	root.Left = build(nums, start, maxIdx)
	root.Right = build(nums, maxIdx+1, end)
	return root
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
