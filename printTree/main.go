package main

import (
	"fmt"
	"strconv"
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

func printTree(root *TreeNode) [][]string {
	h := getHeight(root, 0)
	nodeCnt := 1
	for i := 0; i < h; i++ {
		nodeCnt *= 2
	}
	nodeCnt--
	ans := make([][]string, h)
	for i := range h {
		ans[i] = make([]string, nodeCnt)
	}
	dfs(root, ans, 0, 0, nodeCnt-1)
	return ans
}

func dfs(root *TreeNode, ans [][]string, h, l, r int) {
	if root == nil {
		return
	}
	mid := (l + r) / 2
	ans[h][mid] = strconv.Itoa(root.Val)
	if root.Left != nil {
		dfs(root.Left, ans, h+1, l, mid-1)
	}
	if root.Right != nil {
		dfs(root.Right, ans, h+1, mid+1, r)
	}
}

func getHeight(node *TreeNode, h int) int {
	if node == nil {
		return h
	}

	leftH := getHeight(node.Left, h+1)
	rightH := getHeight(node.Right, h+1)
	return max(leftH, rightH)
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	fmt.Println(printTree(tree))
}
