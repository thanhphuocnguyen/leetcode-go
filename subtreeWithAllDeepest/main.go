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

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	ans, _ := dfs(root, 0)
	return ans
}

func dfs(node *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return nil, depth
	}
	leftNode, leftHeight := dfs(node.Left, depth+1)
	rightNode, rightHeight := dfs(node.Right, depth+1)
	if leftHeight > rightHeight {
		return leftNode, leftHeight
	}
	if rightHeight > leftHeight {
		return rightNode, rightHeight
	}
	return node, leftHeight
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	subtreeWithAllDeepest(root)
}
