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

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var backtrack func(node *TreeNode, path []int, sum int)

	backtrack = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum -= node.Val
		if node.Left == nil && node.Right == nil && sum == 0 {
			cp := make([]int, len(path))
			copy(cp, path)
			ans = append(ans, cp)
			return
		}
		backtrack(node.Left, path, sum)
		backtrack(node.Right, path, sum)

	}
	backtrack(root, []int{}, targetSum)

	return ans
}

func main() {
	// [5,4,8,11,null,13,4,7,2,null,null,5,1]
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}
