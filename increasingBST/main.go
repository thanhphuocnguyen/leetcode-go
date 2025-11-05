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

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(node *TreeNode)
	var newRoot *TreeNode
	curr := newRoot

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if curr == nil {

		}
		if newRoot == nil {
			newRoot = &TreeNode{
				Val: node.Val,
			}
			curr = newRoot
		} else {
			curr.Right = &TreeNode{
				Val: node.Val,
			}
			curr = curr.Right
		}
		dfs(node.Right)
	}

	dfs(root)
	return newRoot
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
	}
	increasingBST(root)
}
