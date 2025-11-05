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
type FindElements struct {
	values map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	values := make(map[int]bool)
	var recover func(root *TreeNode)
	recover = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			root.Left.Val = root.Val*2 + 1
			values[root.Left.Val] = true
			recover(root.Left)
		}
		if root.Right != nil {
			root.Right.Val = root.Val*2 + 2
			values[root.Right.Val] = true
			recover(root.Right)
		}

	}
	recover(root)
	return FindElements{
		values,
	}
}

func (this *FindElements) Find(target int) bool {
	return this.values[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
