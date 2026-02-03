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

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(averageOfLevels(root))
}
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	cnt := make([]int, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > len(ans) {
			ans = append(ans, float64(node.Val))
			cnt = append(cnt, 1)
		} else {
			cnt[level-1]++
			ans[level-1] += float64(node.Val)
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 1)
	for i := range ans {
		ans[i] /= float64(cnt[i])
	}
	return ans
}
