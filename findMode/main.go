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

func findMode(root *TreeNode) []int {
	ans := []int{root.Val}
	var dfs func(node *TreeNode)
	maxCnt := 0
	currVal, currCnt := 0, 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if maxCnt == 0 {
			currVal = node.Val
		}
		if node.Val == currVal {
			currCnt++
		} else {
			currVal = node.Val
			currCnt = 1
		}
		if currCnt > maxCnt {
			maxCnt = currCnt
			ans = []int{node.Val}
		} else if currCnt == maxCnt && node.Val != ans[len(ans)-1] {
			ans = append(ans, node.Val)
		}
		dfs(node.Right)
	}

	dfs(root)

	return ans
}

func main() {
	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println(findMode(root2))
	root1 := &TreeNode{
		Val: 1,
	}
	fmt.Println(findMode(root1))
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(findMode(root))

}
