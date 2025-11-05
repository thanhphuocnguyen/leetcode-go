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

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	vals1, vals2 := make([]int, 0), make([]int, 0)
	dfs(root1, &vals1)
	dfs(root2, &vals2)
	i, j := 0, 0
	ans := make([]int, len(vals1)+len(vals2))
	k := 0
	for i < len(vals1) && j < len(vals2) {
		if vals1[i] < vals2[j] {
			ans[k] = vals1[i]
			i++

		} else {
			ans[k] = vals2[j]
			j++
		}
		k++
	}
	for i < len(vals1) {
		ans[k] = vals1[i]
		i++
		k++
	}
	for j < len(vals2) {
		ans[k] = vals2[j]
		j++
		k++
	}
	return ans
}

func dfs(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, vals)
	*vals = append(*vals, root.Val)
	dfs(root.Right, vals)
}

func main() {
	root1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	getAllElements(root1, root2)
}
