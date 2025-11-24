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
type BSTIterator struct {
	pointer int
	vals    []int
}

func Constructor(root *TreeNode) BSTIterator {
	vals := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return BSTIterator{
		pointer: 0,
		vals:    vals,
	}
}

func (this *BSTIterator) Next() int {
	ans := this.vals[this.pointer]
	this.pointer++
	return ans
}

func (this *BSTIterator) HasNext() bool {
	return this.pointer < len(this.vals)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
