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

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{
		Val: preorder[0],
	}
	st := []*TreeNode{root}
	for i := 1; i < len(preorder); i++ {
		val := preorder[i]
		peek := st[len(st)-1]
		newNode := &TreeNode{Val: preorder[i]}
		if val < peek.Val {
			peek.Left = newNode
		} else {
			curr := st[len(st)-1]
			for len(st) > 0 && st[len(st)-1].Val < val {
				curr = st[len(st)-1]
				st = st[:len(st)-1]
			}
			curr.Right = newNode
		}
		st = append(st, newNode)
	}
	return root
}
func main() {
	fmt.Println(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
}
