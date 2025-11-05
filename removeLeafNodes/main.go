package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = removeLeafNodes(root.Left, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

// func dfs(node *TreeNode, target int) bool {
// 	if node == nil {
// 		return false
// 	}

// 	if node.Left == nil && node.Right == nil && node.Val == target {
// 		return true
// 	}

// 	if dfs(node.Left, target) {
// 		node.Left = nil
// 	}
// 	if dfs(node.Right, target) {
// 		node.Right = nil
// 	}
// 	return false
// }

func main() {
	//root = [1,2,3,2,null,2,4], target = 2
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 2,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val:  3,
	// 		Left: &TreeNode{Val: 2},
	// 		Right: &TreeNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	rs := removeLeafNodes(root, 1)
	fmt.Println(rs)
}
