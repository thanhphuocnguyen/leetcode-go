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

func findSecondMinimumValue(root *TreeNode) int {
	firstMin, secondMin := math.MaxInt32, math.MaxInt32
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if firstMin == -1 || root.Val < firstMin {
			secondMin = firstMin
			firstMin = root.Val
		} else if secondMin == -1 || root.Val < secondMin && root.Val != firstMin {
			secondMin = root.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	if firstMin == secondMin || secondMin == math.MaxInt32 {
		return -1
	}

	return secondMin
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(findSecondMinimumValue(root))
}
