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

func maxLevelSum(root *TreeNode) int {
	h := getHeight(root)
	sumArr := make([]int, h+1)
	dfs(root, 0, sumArr)
	maxi := sumArr[0]
	ans := 1
	for i := 1; i < len(sumArr); i++ {
		if sumArr[i] > maxi {
			maxi = sumArr[i]
			ans = i + 1
		}
	}
	return ans
}

func dfs(root *TreeNode, level int, sumArr []int) {
	if root == nil {
		return
	}

	sumArr[level] += root.Val
	dfs(root.Left, level+1, sumArr)
	dfs(root.Right, level+1, sumArr)
}

func getHeight(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	lh := 0
	if root.Left != nil {
		lh = getHeight(root.Left)
	}
	rh := 0
	if root.Right != nil {
		rh = getHeight(root.Right)
	}

	return max(lh, rh) + 1
}

func main() {
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 7,
	// 		Left: &TreeNode{
	// 			Val: 7,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: -8,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 0,
	// 	},
	// }
	root := &TreeNode{
		Val: 989,
		// Left: &TreeNode{},
		Right: &TreeNode{
			Val: 10250,
			Left: &TreeNode{
				Val: 98693,
			},
			Right: &TreeNode{
				Val: -89388,
				Right: &TreeNode{
					Val: -32127,
				},
			},
		},
	}
	fmt.Println(maxLevelSum(root))
}
