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

func buildTree(preorder []int, inorder []int) *TreeNode {
	var prorderIdx = 0
	root := dfs(preorder, inorder, &prorderIdx, 0, len(inorder)-1)
	return root
}

func dfs(preorder []int, inorder []int, preorderIdx *int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	val := preorder[*preorderIdx]
	*preorderIdx++
	node := TreeNode{val, nil, nil}
	inorderIdx := searchIdx(inorder, val, left, right)
	if inorderIdx < 0 {
		return nil
	}
	node.Left = dfs(preorder, inorder, preorderIdx, left, inorderIdx-1)
	node.Right = dfs(preorder, inorder, preorderIdx, inorderIdx+1, right)
	return &node
}

func searchIdx(inorder []int, target, left, right int) int {
	for i := left; i <= right; i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
