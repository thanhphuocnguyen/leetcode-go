package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var dfs func(node *TreeNode)
	minVal := math.MaxInt32
	prev := -1
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != -1 {
			minVal = min(minVal, node.Val-prev)
		}
		prev = node.Val
		dfs(node.Right)
	}

	dfs(root)
	return minVal
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
	fmt.Println(getMinimumDifference(root))
}
