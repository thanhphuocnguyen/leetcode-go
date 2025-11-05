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

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ans, _ := postOrderTraversal(root, 0)
	return ans
}

func postOrderTraversal(node *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return nil, depth
	}
	leftLca, lHeight := postOrderTraversal(node.Left, depth+1)
	rightLca, rHeight := postOrderTraversal(node.Right, depth+1)

	if lHeight > rHeight {
		return leftLca, lHeight
	}
	if rHeight > lHeight {
		return rightLca, rHeight
	}
	return leftLca, lHeight
}

func main() {
	// root [3,5,1,6,2,0,8,null,null,7,4]
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
	lcaDeepestLeaves(root)
}
